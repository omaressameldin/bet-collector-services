package document

import (
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"

	"net/http"
)

func Handler(client *elasticsearch7.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPut:
		bulkAddDocuments(client, w, r)
		return
	case http.MethodHead:
		return
	case http.MethodGet:
		queryDocuments(client, w, r)
		return
	}

	http.Error(w, "", http.StatusNotFound)
}

func NameHandler(client *elasticsearch7.Client, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDocument(client, w, r)
		return
	case http.MethodDelete:
		return
	}
	http.Error(w, "", http.StatusNotFound)
}
