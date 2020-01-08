package document

import (
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/internal/helpers"
)

const (
	searchIndexKey = "indices"
)

func queryDocuments(
	client *elasticsearch7.Client,
	w http.ResponseWriter,
	r *http.Request,
) {
	req := esapi.SearchRequest{
		Index: r.URL.Query()[searchIndexKey],
		Body:  r.Body,
	}

	res, err, statusCode := helpers.SendRequest(client, req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Write([]byte(*res))
}
