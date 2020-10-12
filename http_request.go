package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

//Todo chnage file name

//Todo change struct name
type httpRequest struct {
	httpClient HTTPClient
}

//Todo change func name
func (dc *httpRequest) createRequest(targetUrl string, body []map[string]interface{}) []error {
	var error []error
	for _, value := range body {
		reqBody, err := parseBody(value)
		if err != nil {
			error = append(error, err)
			continue
		}

		response, err := dc.httpClient.Post(targetUrl, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			error = append(error, err)
			continue
		}

		if code := response.StatusCode; code >= 400 && code <= 600 {
			buf := new(strings.Builder)
			_, err = io.Copy(buf, response.Body)

			error = append(error, fmt.Errorf("unable to seed %s because %s", reqBody, buf.String()))

		}
	}
	return error
}

func parseBody(body map[string]interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func NewHTTPRequest(client HTTPClient) *httpRequest {
	return &httpRequest{
		httpClient: client,
	}
}
