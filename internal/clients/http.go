package clients

import (
	"net/http"
	"time"
)

func NewDefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
