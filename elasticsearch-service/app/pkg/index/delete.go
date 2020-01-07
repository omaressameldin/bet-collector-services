package index

import (
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gorilla/mux"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/internal/helpers"
)

func deleteIndex(
	client *elasticsearch7.Client,
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	indexName := vars["name"]
	req := esapi.IndicesDeleteRequest{
		Index: []string{indexName},
	}

	res, err, statusCode := helpers.SendRequest(client, req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Write([]byte(*res))
}
