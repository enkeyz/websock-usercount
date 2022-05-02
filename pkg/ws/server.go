package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type UserCountServer struct {
	http.Handler
	hub *Hub
}

func NewUserCountServer(hub *Hub) *UserCountServer {
	ucs := &UserCountServer{}

	ucs.hub = hub

	router := http.NewServeMux()
	router.Handle("/usercount", http.HandlerFunc(ucs.wsHandler))

	ucs.Handler = router

	return ucs
}

func (u *UserCountServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 2048,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error when upgrading connection %v", err)
		return
	}

	client := NewClient(conn, u.hub)
	u.hub.join <- client
	go client.Listen()
}
