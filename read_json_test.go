package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadJson(t *testing.T) {
	testJsonFilePath := "./test.json"
	tests := []struct {
		name           string
		actualResult   func() ([]map[string]interface{}, error)
		expectedResult []map[string]interface{}
		expectedError  error
	}{
		{
			name: "reading json successfully",
			actualResult: func() ([]map[string]interface{}, error) {
				return readJson(testJsonFilePath)
			},
			expectedResult: []map[string]interface{}{{"description": "this is to test json", "name": "test"}},
			expectedError:  nil,
		},
		{
			name: "invalid json path",
			actualResult: func() ([]map[string]interface{}, error) {
				return readJson("")
			},
			expectedResult: nil,
			expectedError:  errors.New("open : no such file or directory"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result, err := test.actualResult()
			assert.Equal(tt, test.expectedResult, result)
			if err != nil {
				assert.Equal(tt, test.expectedError.Error(), err.Error())

			}
		})
	}
}
