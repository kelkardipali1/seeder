package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
		reqBody, err:= parseBody(value)
		if err != nil {
			return err
		}
		fmt.Println("--->",string(reqBody))
		//req, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer(reqBody))
		//if err != nil {
		//	return err
		//}
		//req.Header.Add("Content-Type", "application/json")

	}

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
