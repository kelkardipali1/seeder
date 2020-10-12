package main

import (
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type httpClient struct {
	hc *http.Client
}

func (client *httpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return client.hc.Post(url, contentType, body)
}

func NewHttpClient() HTTPClient {
	return &httpClient{
		hc: &http.Client{
			Timeout: time.Second * time.Duration(10),
		},
	}
}
