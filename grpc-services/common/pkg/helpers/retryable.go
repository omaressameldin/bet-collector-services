package helpers

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	retryTime time.Duration = 5 * time.Second
)

//CreateRetryableClient creates an httpClient that retries requests 5 times before failing
func CreateRetryableClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryWaitMin = retryTime

	return client
}
