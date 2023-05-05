package main_test

import (
	"practice/api"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetKey(t *testing.T) {
	testCases := []struct {
		name         string
		key          string
		keyValue     map[string]string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "get map data by key",
			key:          "foo",
			keyValue:     map[string]string{"foo": "bar"},
			expectedCode: 200,
			expectedBody: "bar",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ginCtx := &gin.Context{}
			output := api.GetKeyAPI(ginCtx, tc.key, tc.keyValue)
			assert.Equal(t, tc.expectedBody, output)
		})
	}
}