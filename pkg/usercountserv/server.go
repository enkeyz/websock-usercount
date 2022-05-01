package usercountserv

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type UserCountServer struct {
	http.Handler
}

func NewUserCountServer() *UserCountServer {
	ucs := &UserCountServer{}

	router := http.NewServeMux()
	router.Handle("/usercount", http.HandlerFunc(ucs.userCountHandler))

	ucs.Handler = router

	return ucs
}

func (u *UserCountServer) userCountHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 4096,
	}
	// TODO never do this
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error when upgrading connection %v", err)
		return
	}
	log.Printf("Successfull websocket connection from %q", conn.RemoteAddr().String())
}
