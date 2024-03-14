package core

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

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
	clientsMute   *sync.Mutex
	httpServer    *http.Server
	serveMux      http.ServeMux
}

func NewCore(address string, channels uint64, busses uint64, logFile *log.Logger) *Core {
	c := Core{
		RunningConfig: *config.NewShowConfig("NewShow", "NewShow.cfg", channels, busses),
		logFile:       logFile,
		address:       address,
		stop:          make(chan bool),
		rxChannel:     make(chan api.Command),
		clients:       make(map[uuid.UUID]*api.Client),
		clientsMute:   &sync.Mutex{},
	}

	// Setup serve mux
	c.serveMux.HandleFunc("/api/v1/connect/", c.apiV1Handler)
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
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
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
		cfg, _ := json.Marshal(c.RunningConfig)
		c.notifyClients(api.Request{
			Method: api.SHOW_LOAD,
			Path:   msg.RequestData.Path,
			Data:   string(cfg),
		})
	}
}

func (c *Core) notifyClients(msg api.Request) {
	c.clientsMute.Lock()
	defer c.clientsMute.Unlock()

	for id, cl := range c.clients {
		if !cl.Active {
			delete(c.clients, id)
			continue
		}
		cl.TxChannel <- msg
	}
}

func (c *Core) LoadShowConfig(filepath string) error {
	cfgBytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	cfg := config.ShowConfig{}
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		return err
	}

	c.RunningConfig = cfg

	msg := api.Request{
		Method: api.SHOW_LOAD,
		Path:   "",
		Data:   string(cfgBytes[:]),
	}
	c.notifyClients(msg)

	return nil
}

func (c *Core) SaveShowConfig(filepath string) error {
	cfg, err := json.Marshal(c.RunningConfig)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, cfg, 0644)
	if err != nil {
		return err
	}

	return nil
}
