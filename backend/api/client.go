package api

import (
	"context"
	"log"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Handles the websocket connection to the client
type Client struct {
	UUID      uuid.UUID
	conn      *websocket.Conn
	Active    bool
	ConnCtx   context.Context
	logFile   *log.Logger
	TxChannel chan Request
	rxChannel chan Command
	stopSend  chan bool
}

func NewClient(conn *websocket.Conn, id uuid.UUID, rxChannel chan Command, logFile *log.Logger) *Client {
	c := Client{
		UUID:      id,
		conn:      conn,
		Active:    true,
		ConnCtx:   context.Background(),
		logFile:   logFile,
		TxChannel: make(chan Request),
		rxChannel: rxChannel,
		stopSend:  make(chan bool),
	}
	return &c
}

func (c *Client) Run() {
	go c.receiveService()
	go c.sendService()
}

func (c *Client) receiveService() {
	for {
		var v Request
		err := wsjson.Read(c.ConnCtx, c.conn, &v)
		if err != nil {
			c.logFile.Printf("Client %v encountered error %v", c.UUID, err)
			c.conn.Close(websocket.StatusGoingAway, "")
			c.stopSend <- true
			c.Active = false
			return
		}
		c.rxChannel <- Command{
			ClientID:    c.UUID,
			RequestData: v,
		}
	}
}

func (c *Client) sendService() {
	for {
		select {
		case <-c.stopSend:
			return
		case cmd := <-c.TxChannel:
			wsjson.Write(c.ConnCtx, c.conn, cmd)
		}
	}
}
