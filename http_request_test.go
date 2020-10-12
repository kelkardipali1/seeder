package main

import (
	errors2 "errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHTTPClient struct{
	statusCode int
	err error
}

func (hc *mockHTTPClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	res := httptest.NewRecorder()
	res.WriteHeader(hc.statusCode)
	return res.Result(), hc.err
}

func TestCreateRequest(t *testing.T) {
	tests := []struct {
		name          string
		actualError   func() []error
		expectedError []error
	}{
		{
			name: "request successful",
			actualError: func() []error {
				hc := &mockHTTPClient{
					statusCode: http.StatusCreated,
					err:        nil,
				}
				body := []map[string]interface{}{{"key": "value"}, {"key": "value"}}
				req := NewHTTPRequest(hc)
				errors := req.createRequest("/url", body)
				return errors
			},
			expectedError: nil,
		},
		{
			name: "request failure",
			actualError: func() []error {
				hc := &mockHTTPClient{
					statusCode: http.StatusConflict,
					err:        errors2.New("record already exist"),
				}
				body := []map[string]interface{}{{"key": "value"}, {"key": "value"}}
				req := NewHTTPRequest(hc)
				errors := req.createRequest("/url", body)
				return errors
			},
			expectedError: []error{errors2.New("record already exist"),errors2.New("record already exist")},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(t, len(test.expectedError), len(test.actualError()))
		})
	}
}
