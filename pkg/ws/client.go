package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func NewClient(connection *websocket.Conn) *Client {
	return &Client{
		conn: connection,
	}
}

func (c *Client) Send(activeConnection int) {
	log.Printf("Sending client: number of active clients: %d", activeConnection)
}
