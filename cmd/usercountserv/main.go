package main

import (
	"log"
	"net/http"

	"github.com/enkeyz/websock-usercount/pkg/usercountserv"
)

func main() {
	log.Println("Websocket listening on http://127.0.0.0:7777...")
	log.Fatal(http.ListenAndServe(":7777", usercountserv.NewUserCountServer()))
}
