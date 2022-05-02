package ws

import (
	"time"
)

const tickerDuration = 1 * time.Second

type Hub struct {
	clients map[*Client]bool
	ticker  *time.Ticker
	join    chan *Client
	leave   chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
		ticker:  time.NewTicker(tickerDuration),
		join:    make(chan *Client),
		leave:   make(chan *Client),
	}
}

func (h *Hub) activeConnections() int {
	return len(h.clients)
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.join:
			h.clients[client] = true
		case client := <-h.leave:
			delete(h.clients, client)
		case <-h.ticker.C:
			for client := range h.clients {
				select {
				case client.send <- h.activeConnections():
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
