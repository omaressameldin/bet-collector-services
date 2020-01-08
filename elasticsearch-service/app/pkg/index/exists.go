package index

import (
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/internal/helpers"
)

const existsQueryKey = "indices"

func doesIndexExist(
	client *elasticsearch7.Client,
	w http.ResponseWriter,
	r *http.Request,
) {
	indices := r.URL.Query()[existsQueryKey]

	req := esapi.IndicesExistsRequest{
		Index: indices,
	}
	res, err, statusCode := helpers.SendRequest(client, req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	w.Write([]byte(*res))
}
