package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDoSomething(t *testing.T){
	tests := []struct {
		name           string
		actualResult   func()
	}{
		{
			name: "sample"
			actualResult: func() {
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, , )
		})
	}
}

