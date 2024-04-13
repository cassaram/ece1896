package backend

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/cassaram/ece1896/backend/config"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type BackendConnection struct {
	address      string
	config       config.ShowConfig
	websocket    *websocket.Conn
	websocketCtx context.Context
	logFile      *log.Logger
	txChannel    chan Request
	subscribers  []Subscriber
}

func NewBackendConnection(address string, log *log.Logger) *BackendConnection {
	b := BackendConnection{
		address:     address,
		logFile:     log,
		txChannel:   make(chan Request),
		subscribers: make([]Subscriber, 0),
	}
	return &b
}

func (b *BackendConnection) Connect() error {
	b.websocketCtx = context.Background()
	ws, _, err := websocket.Dial(b.websocketCtx, b.address, nil)
	if err != nil {
		return err
	}
	b.websocket = ws
	go b.rxHandler()
	go b.txHandler()
	return nil
}

func (b *BackendConnection) Subscribe(sub Subscriber) {
	b.subscribers = append(b.subscribers, sub)
}

func (b *BackendConnection) Send(msg Request) {
	b.txChannel <- msg
}

func (b *BackendConnection) GetConfig() config.ShowConfig {
	return b.config
}

func (b *BackendConnection) rxHandler() {
	for {
		var v Request
		err := wsjson.Read(b.websocketCtx, b.websocket, &v)
		if err != nil {
			b.logFile.Printf("Error in backend connection: %v\n", err.Error())
			continue
		}
		b.handleRequest(v)
	}
}

func (b *BackendConnection) txHandler() {
	for {
		v := <-b.txChannel
		err := wsjson.Write(b.websocketCtx, b.websocket, v)
		b.logFile.Println("Writing", v)
		if err != nil {
			b.logFile.Printf("Error in backend connection: %v\n", err.Error())
			continue
		}
	}
}

func (b *BackendConnection) handleRequest(v Request) {
	b.logFile.Printf("Received: %v\n", v)
	switch v.Method {
	case ERROR:
		b.logFile.Printf("Backend returned error %s\n", v.Data)
	case SHOW_GET:
		paths := strings.Split(v.Path, ".")
		err := b.config.SetValue(paths, v.Data)
		if err != nil {
			b.logFile.Printf("Error setting value (%s) at path %s: %s", v.Data, v.Path, err.Error())
		}
		for _, sub := range b.subscribers {
			sub.UpdatePath(v.Path, b.config)
		}
	case SHOW_SET:
		paths := strings.Split(v.Path, ".")
		err := b.config.SetValue(paths, v.Data)
		if err != nil {
			b.logFile.Printf("Error setting value (%s) at path %s: %s", v.Data, v.Path, err.Error())
		}
		for _, sub := range b.subscribers {
			sub.UpdatePath(v.Path, b.config)
		}
	case SHOW_LOAD:
		var cfg config.ShowConfig
		err := json.Unmarshal([]byte(v.Data), &cfg)
		if err != nil {
			b.logFile.Printf("Error parsing show: %s", err.Error())
		}
		b.config = cfg
		for _, sub := range b.subscribers {
			sub.ReloadConfig(cfg)
		}
	}
}
