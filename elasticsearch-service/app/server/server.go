package server

import (
	"fmt"
	"log"
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/pkg/index"
)

func RunServer(client *elasticsearch7.Client, port int) error {
	r := mux.NewRouter()

	r.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		index.Handler(client, w, r)
	})

	r.HandleFunc("/index/{name}", func(w http.ResponseWriter, r *http.Request) {
		index.NameHandler(client, w, r)
	})
	log.Printf("starting server on port: %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
