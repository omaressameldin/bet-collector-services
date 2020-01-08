package document

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/internal/helpers"
)

type Document struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

type BulkAddRequest struct {
	Index     string     `json:"index"`
	Documents []Document `json:"documents"`
}

func createBulkBody(documents []Document) ([]byte, error) {
	var refactoredData bytes.Buffer
	for _, document := range documents {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, document.ID, "\n"))

		data, err := json.Marshal(document.Data)
		if err != nil {
			return nil, err
		}
		refactoredData.Grow(len(meta) + len(data) + len([]byte("\n")))
		refactoredData.Write(meta)
		refactoredData.Write(data)
		refactoredData.Write([]byte("\n"))
	}

	return refactoredData.Bytes(), nil
}

func bulkAddDocuments(
	client *elasticsearch7.Client,
	w http.ResponseWriter,
	r *http.Request,
) {
	var bulkRequest BulkAddRequest
	err := json.NewDecoder(r.Body).Decode(&bulkRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := createBulkBody(bulkRequest.Documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	req := esapi.BulkRequest{
		Index: bulkRequest.Index,
		Body:  bytes.NewReader(body),
	}
	res, err, statusCode := helpers.SendRequest(client, req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.Write([]byte(*res))
}
