package main

import (
	"net/http"
	"time"
)

type HTTPClient interface {

}

type defaultHTTPClient struct {
	client *http.Client
	target string
	body   []byte
}


func NewHttpClient(target string,body []byte) HTTPClient {
	return &defaultHTTPClient{
		client: &http.Client{
			Timeout: time.Second * time.Duration(10),
		},
		target: target,
		body:   body,
	}
}

