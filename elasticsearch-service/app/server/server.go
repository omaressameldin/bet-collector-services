package server

import (
	"fmt"
	"log"
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to do here"))
}

func RunServer(client *elasticsearch7.Client, port int) error {
	http.HandleFunc("/", handler)
	log.Println("starting server on port: %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
