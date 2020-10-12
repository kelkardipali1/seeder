package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strings"
)

type HTTPRequest interface {
	createRequest(targetUrl string, body []map[string]interface{})
}

type httpRequest struct {
	httpClient HTTPClient
}

func (dc *httpRequest) createRequest(targetUrl string, body []map[string]interface{}) {
	for _, value := range body {
		reqBody, err := parseBody(value)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		response, err := dc.httpClient.Post(targetUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if code := response.StatusCode; code >= 400 && code <= 600 {
			buf := new(strings.Builder)
			_, err = io.Copy(buf, response.Body)

			log.Printf("unable to seed %s because %s", reqBody, buf.String())
		}
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

func NewHTTPRequest() HTTPRequest {
	return &httpRequest{
		httpClient: NewHttpClient(),
	}
}
