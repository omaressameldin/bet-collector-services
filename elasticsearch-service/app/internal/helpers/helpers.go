package helpers

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func SendRequest(client *elasticsearch7.Client, req esapi.Request) (*string, error, int) {
	res, err := req.Do(context.Background(), client)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	resString := buf.String()
	if res.IsError() {
		return nil, errors.New(resString), res.StatusCode
	}

	return &resString, nil, res.StatusCode
}
