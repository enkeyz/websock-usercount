package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enkeyz/websock-usercount/pkg/usercountserv"
)

const port = 7777

func main() {
	log.Println("Websocket listening on http://127.0.0.0:7777...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), usercountserv.NewUserCountServer()))
}
