package ws

import (
	"sync"
	"time"
)

const tickerDuration = 1 * time.Second

type Hub struct {
	mu          sync.Mutex
	connections map[*Client]bool
	ticker      *time.Ticker
}

func NewHub() *Hub {
	return &Hub{
		mu:          sync.Mutex{},
		connections: make(map[*Client]bool),
		ticker:      time.NewTicker(tickerDuration),
	}
}

func (h *Hub) Join(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.connections[client] = true
}

func (h *Hub) Leave(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.connections, client)
}

func (h *Hub) ActiveConnections() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.connections)
}

func (h *Hub) Run() {
	for {
		select {
		case <-h.ticker.C:
			for client := range h.connections {
				client.Send(h.ActiveConnections())
			}
		}
	}
}
