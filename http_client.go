package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type HTTPClient interface {
	doSomething(targetUrl string, body []map[string]interface{}) error
}

type defaultHTTPClient struct {
	client *http.Client
}

func (dc *defaultHTTPClient) doSomething(targetUrl string, body []map[string]interface{}) error {
	for _, value := range body {
		reqBody, err := parseBody(value)
		if err != nil {
			log.Fatal("Json can not be parse")
			return err
		}

		response, err := dc.client.Post(targetUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatal("Invalid url")
			return err
		}

		buf := new(strings.Builder)
		_, err = io.Copy(buf, response.Body)
		log.Println("response", buf.String())

	}
	return nil
}

func parseBody(body map[string]interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func NewHttpClient() HTTPClient {
	return &defaultHTTPClient{
		client: &http.Client{
			Timeout: time.Second * time.Duration(10),
		},
	}
}
