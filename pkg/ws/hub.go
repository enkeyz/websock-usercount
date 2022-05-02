package ws

import (
	"log"
	"time"
)

const tickerDuration = 3 * time.Second

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
			log.Printf("Successfull websocket connection from %s", client.remoteAddr)
		case client := <-h.leave:
			log.Printf("Closing websocket connection on %s", client.remoteAddr)
			delete(h.clients, client)
		case <-h.ticker.C:
			if h.activeConnections() == 0 {
				continue
			}

			log.Printf("Broadcasting message to %d clients", h.activeConnections())
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
