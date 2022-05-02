package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enkeyz/websock-usercount/pkg/ws"
)

const port = 7777

func main() {
	hub := ws.NewHub()
	go hub.Run()
	log.Printf("Websocket listening on http://127.0.0.1:%d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), ws.NewUserCountServer(hub)))
}
