package index

import (
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"

	"net/http"
)

func Handler(client *elasticsearch7.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPut:
		addIndex(client, w, r)
		return
	case http.MethodHead:
		doesIndexExist(client, w, r)
		return
	}

	http.Error(w, "", http.StatusNotFound)
}

func NameHandler(client *elasticsearch7.Client, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getIndex(client, w, r)
		return
	case http.MethodDelete:
		deleteIndex(client, w, r)
		return
	}
	http.Error(w, "", http.StatusNotFound)
}
