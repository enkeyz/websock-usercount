package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/", fs)

	log.Println("Webserver listening on http://127.0.0.0:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
