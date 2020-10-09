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
	doSomething(targetUrl string, body []map[string]interface{})
}

type defaultHTTPClient struct {
	client *http.Client
}

func (dc *defaultHTTPClient) doSomething(targetUrl string, body []map[string]interface{}) {
	for _, value := range body {
		reqBody, err := parseBody(value)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		response, err := dc.client.Post(targetUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if code := response.StatusCode; code >= 400 && code <= 600 {
			buf := new(strings.Builder)
			_, err = io.Copy(buf, response.Body)

			log.Printf("unable to seed %s because %s", reqBody, buf.String())
		}
		//else if code := response.StatusCode; code >= 200 && code < 300 {
		//	log.Println("data seed successfully")
		//}

	}
	log.Println("done")

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
