package api

import (
	"context"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Handles the websocket connection to the client
type Client struct {
	conn    *websocket.Conn
	connCtx context.Context
}

func (c *Client) receiveService() {
	var v interface{}
	wsjson.Read(c.connCtx, c.conn, &v)
}
