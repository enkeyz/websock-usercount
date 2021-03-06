package ws

import (
	"strconv"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	remoteAddr string
	send       chan int
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	return &Client{
		hub:        hub,
		conn:       connection,
		remoteAddr: connection.RemoteAddr().String(),
		send:       make(chan int),
	}
}

func (c *Client) Listen() {
	defer func() {
		c.hub.leave <- c
		c.conn.Close()
	}()

	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(msg)))
		if err != nil {
			return
		}
	}
}
