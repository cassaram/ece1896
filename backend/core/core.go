package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/cassaram/ece1896/backend/api"
	"github.com/cassaram/ece1896/backend/config"
	"github.com/google/uuid"
	"nhooyr.io/websocket"
)

type Core struct {
	RunningConfig config.ShowConfig
	logFile       *log.Logger
	address       string
	stop          chan bool
	rxChannel     chan api.Command
	clients       map[uuid.UUID]*api.Client
	clientsMute   sync.Mutex
	httpServer    *http.Server
	serveMux      http.ServeMux
	cfgPath       string
}

func NewCore(address string, channels uint64, busses uint64, logFile *log.Logger, cfgPath string) *Core {
	// Find a valid config
	c := Core{
		//RunningConfig: *config.NewShowConfig("NewShow", "NewShow.cfg", channels, busses),
		logFile:   logFile,
		address:   address,
		stop:      make(chan bool),
		rxChannel: make(chan api.Command),
		clients:   make(map[uuid.UUID]*api.Client),
		cfgPath:   cfgPath,
	}

	// Find latest config
	configs, err := c.GetShowConfigs()
	loadedConfig := false
	if err != nil || len(configs) == 0 {
		c.RunningConfig = *config.NewShowConfig("NewShow", "NewShow.showcfg", channels, busses)
		c.logFile.Printf("No configs found. Loading new show file.")
		loadedConfig = true
	}
	for _, cfg := range configs {
		if cfg.FileName == "LATEST.showcfg" {
			if err := c.LoadShowConfig(cfg.FileName); err != nil {
				c.RunningConfig = *config.NewShowConfig("NewShow", "NewShow.showcfg", channels, busses)
				c.logFile.Printf("LATEST.showcfg could not be loaded. Loading new show file. Error: %v", err)
			}
			loadedConfig = true
		}
	}
	if !loadedConfig {
		c.RunningConfig = *config.NewShowConfig("NewShow", "NewShow.showcfg", channels, busses)
		c.logFile.Printf("LATEST.showcfg not found. Loading new show file.")
	}

	// Setup serve mux
	c.serveMux.HandleFunc("/api/v1/ws", c.apiV1Handler)
	c.httpServer = &http.Server{
		Handler:      &c,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return &c
}

func (c *Core) Run() {
	// Start http server
	listener, err := net.Listen("tcp", c.address)
	if err != nil {
		log.Println(err)
	}

	go func() {
		c.httpServer.Serve(listener)
	}()

	c.logFile.Printf("Backend app running on tcp %v", c.address)

	// Handle messages
	for {
		select {
		case msg := <-c.rxChannel:
			c.handleMessage(msg)
		case <-c.stop:
			return
		}
	}
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.serveMux.ServeHTTP(w, r)
}

func (c *Core) apiV1Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		c.logFile.Println(err)
		return
	}
	client := api.NewClient(conn, uuid.New(), c.rxChannel, c.logFile)
	client.Run()
	c.clientsMute.Lock()
	defer c.clientsMute.Unlock()
	c.clients[client.UUID] = client

	// Send it the initial data via full-config
	cfg, _ := json.Marshal(c.RunningConfig)
	client.TxChannel <- api.Request{
		Method: api.SHOW_LOAD,
		Path:   "",
		Data:   string(cfg),
	}
}

func (c *Core) handleMessage(msg api.Command) {
	switch msg.RequestData.Method {
	case api.SHOW_GET:
		path := strings.Split(msg.RequestData.Path, ".")
		val, err := c.RunningConfig.GetValue(path)
		c.clientsMute.Lock()
		defer c.clientsMute.Unlock()
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		c.clients[msg.ClientID].TxChannel <- api.Request{
			Method: api.SHOW_GET,
			Path:   msg.RequestData.Path,
			Data:   val,
		}
	case api.SHOW_SET:
		path := strings.Split(msg.RequestData.Path, ".")
		err := c.RunningConfig.SetValue(path, msg.RequestData.Data)
		c.clientsMute.Lock()
		defer c.clientsMute.Unlock()
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		c.notifyClients(api.Request{
			Method: api.SHOW_SET,
			Path:   msg.RequestData.Path,
			Data:   msg.RequestData.Data,
		})
		c.SaveCurrentShowConfig()
	case api.SHOW_LOAD:
		err := c.LoadShowConfig(msg.RequestData.Path)
		c.clientsMute.Lock()
		defer c.clientsMute.Unlock()
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		cfg, err := json.Marshal(c.RunningConfig)
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		c.notifyClients(api.Request{
			Method: api.SHOW_LOAD,
			Path:   msg.RequestData.Path,
			Data:   string(cfg),
		})
		c.SaveCurrentShowConfig()
	case api.SHOW_LIST:
		showConfigs, err := c.GetShowConfigs()
		c.clientsMute.Lock()
		defer c.clientsMute.Unlock()
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		cfgsJson, err := json.Marshal(showConfigs)
		if err != nil {
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   msg.RequestData.Path,
				Data:   err.Error(),
			}
			break
		}
		c.clients[msg.ClientID].TxChannel <- api.Request{
			Method: api.SHOW_LIST,
			Path:   msg.RequestData.Path,
			Data:   string(cfgsJson),
		}
	case api.SHOW_SAVE:
		err := c.SaveShowConfig(c.RunningConfig.FileName, c.RunningConfig)
		if err != nil {
			c.clientsMute.Lock()
			defer c.clientsMute.Unlock()
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   "",
				Data:   err.Error(),
			}
		}
	case api.SHOW_SAVEAS:
		err := c.SaveShowConfig(msg.RequestData.Data, c.RunningConfig)
		if err != nil {
			c.clientsMute.Lock()
			defer c.clientsMute.Unlock()
			c.clients[msg.ClientID].TxChannel <- api.Request{
				Method: api.ERROR,
				Path:   "",
				Data:   err.Error(),
			}
		}
	}
}

func (c *Core) notifyClients(msg api.Request) {
	fmt.Printf("Notifying %d clients of: %v\n", len(c.clients), msg)

	for id, cl := range c.clients {
		if !cl.Active {
			delete(c.clients, id)
			continue
		}
		cl.TxChannel <- msg
	}
}

func (c *Core) LoadShowConfig(filename string) error {
	cfgBytes, err := os.ReadFile(c.cfgPath + "/shows/" + filename)
	if err != nil {
		return err
	}
	cfgCompactBytesBuffer := &bytes.Buffer{}
	err = json.Compact(cfgCompactBytesBuffer, cfgBytes)
	if err != nil {
		return err
	}
	cfgCompactBytes := cfgCompactBytesBuffer.Bytes()

	cfg := config.ShowConfig{}
	err = json.Unmarshal(cfgCompactBytes, &cfg)
	if err != nil {
		return err
	}

	c.RunningConfig = cfg
	c.logFile.Printf("Loading show file: %v", c.cfgPath+"/shows/"+filename)

	msg := api.Request{
		Method: api.SHOW_LOAD,
		Path:   "",
		Data:   string(cfgCompactBytes[:]),
	}
	c.notifyClients(msg)

	return nil
}

func (c *Core) SaveShowConfig(filename string, showCfg config.ShowConfig) error {
	reg, err := regexp.MatchString("\\.showcfg$", filename)
	if err != nil {
		return err
	}
	if !reg {
		return fmt.Errorf("filename does not end in .showcfg: %s", filename)
	}

	showCfg.FileName = filename

	cfg, err := json.MarshalIndent(showCfg, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.cfgPath+"/shows/"+filename, cfg, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) GetShowConfigs() ([]ConfigFile, error) {
	entries, err := os.ReadDir(c.cfgPath + "/shows/")
	if err != nil {
		return []ConfigFile{}, err
	}
	showCfgs := make([]ConfigFile, 0)
	for _, entry := range entries {
		// Don't handle sub directories
		if entry.IsDir() {
			continue
		}
		// Don't handle non-valid files
		info, err := entry.Info()
		if err != nil {
			continue
		}
		// Don't handle symlinks
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			continue
		}
		// Don't handle non *.showcfg files
		if filepath.Ext(entry.Name()) != ".showcfg" {
			continue
		}

		// Read data as little as needed
		data, _ := os.ReadFile(c.cfgPath + "/shows/" + entry.Name())
		dataParsed, _ := jason.NewObjectFromBytes(data)
		showName, _ := dataParsed.GetString("name")

		showCfgs = append(showCfgs, ConfigFile{
			Name:     showName,
			FileName: info.Name(),
			Size:     info.Size(),
			ModTime:  info.ModTime().Format("2006-01-02 15:04:05"),
		})
	}

	return showCfgs, nil
}

func (c *Core) SaveCurrentShowConfig() {
	if err := c.SaveShowConfig("LATEST.showcfg", c.RunningConfig); err != nil {
		c.logFile.Printf("error writing current config: %v", err)
	}
}
