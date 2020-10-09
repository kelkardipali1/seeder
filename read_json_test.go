package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadJson(t *testing.T) {
	tests := []struct {
		name           string
		actualResult   func() ([]map[string]interface{}, error)
		expectedResult []map[string]interface{}
		expectedError  error
	}{
		{
			name: "reading json without array successfully",
			actualResult: func() ([]map[string]interface{}, error) {
				data := []byte(`{"name": "test","description": "this is to test json"}`)
				tempFile := createTempFile(t, data)
				defer os.Remove(tempFile.Name())

				return readJson(tempFile.Name())
			},
			expectedResult: []map[string]interface{}{{"description": "this is to test json", "name": "test"}},
			expectedError:  nil,
		},
		{
			name: "reading json with array successfully",
			actualResult: func() ([]map[string]interface{}, error) {
				data := []byte(`[{"name": "test","description": "this is to test json"},{"name": "test","description": "this is to test json"}]`)
				tempFile := createTempFile(t, data)
				defer os.Remove(tempFile.Name())

				return readJson(tempFile.Name())
			},
			expectedResult: []map[string]interface{}{{"description": "this is to test json", "name": "test"}, {"name": "test", "description": "this is to test json"}},
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

func createTempFile(t *testing.T, contentToWrite []byte) *os.File {
	tempFile, _ := ioutil.TempFile("./", "*test_seed.json")
	_, err := tempFile.Write(contentToWrite)
	require.NoError(t, err)
	return tempFile
}
