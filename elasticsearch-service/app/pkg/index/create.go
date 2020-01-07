package index

import (
	"bytes"
	"encoding/json"
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/internal/helpers"
)

var (
	aggregationMap = map[string]map[string]string{
		"keyword": {
			"type": "keyword",
		},
	}
)

type MappingProperty struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	CanBeAggregated bool   `json:"canBeAggregated"`
}

type Mapping struct {
	Properties []MappingProperty `json:"properties"`
}

type AddIndexRequest struct {
	Name     string  `json:"name"`
	Mappings Mapping `json:"mappings"`
}

func createMappingJson(m Mapping) ([]byte, error) {
	mapping := map[string]map[string]map[string]map[string]interface{}{
		"mappings": {
			"properties": {},
		},
	}

	for _, p := range m.Properties {
		mapping["mappings"]["properties"][p.Name] = make(map[string]interface{})

		mapping["mappings"]["properties"][p.Name]["type"] = p.Type
		if p.CanBeAggregated {
			mapping["mappings"]["properties"][p.Name]["fields"] = aggregationMap
		}
	}

	mappingJSON, err := json.Marshal(mapping)
	if err != nil {
		return nil, err
	}
	return mappingJSON, nil
}

func addIndex(client *elasticsearch7.Client, w http.ResponseWriter, r *http.Request) {
	var addRequest AddIndexRequest
	err := json.NewDecoder(r.Body).Decode(&addRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mappingJSON, err := createMappingJson(addRequest.Mappings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := esapi.IndicesCreateRequest{
		Index: addRequest.Name,
		Body:  bytes.NewReader(mappingJSON),
	}

	res, err, statusCode := helpers.SendRequest(client, req)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(*res))
}
