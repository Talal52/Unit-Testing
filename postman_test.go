package main_test

import (
	"goapi/api"
	"goapi/model"
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
			key:          "talal",
			keyValue:     map[string]string{"talal": "56"},
			expectedCode: 200,
			expectedBody: "56",
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

func TestDeleteKey(t *testing.T) {
	testCases := []struct {
		name         string
		key          string
		keyValue     map[string]string
		expectedCode int
		expectedBody interface{}
	}{
		{
			name:         "delete map data by key",
			key:          "ta",
			keyValue:     map[string]string{"ta": "55"},
			expectedCode: 200,
			expectedBody: map[string]string{},
		},
	}
	for _, ta := range testCases {
		t.Run(ta.name, func(t *testing.T) {
			ginCtx := &gin.Context{}
			output, err := api.DeleteKeyAPI(ginCtx, ta.key, ta.keyValue)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			assert.Equal(t, ta.expectedBody, output)
		})
	}
}
func TestUpdateKey(t *testing.T) {
	//var response model.Response
	testCases := []struct {
		name         string
		key          string
		keyValue     map[string]string
		expectedCode int
		expectedBody map[string]string
	}{
		{
			name:         "update existing key",
			key:          "talal",
			keyValue:     map[string]string{"name": "talal"},
			expectedCode: 200,
			expectedBody: map[string]string{
				"name": "talal",
			},
		},
	}
	for _, st := range testCases {
		t.Run(st.name, func(t *testing.T) {
			// data := st.keyValue
			// req := model.Response{Key: st.key, Value: "newValue"}
			ginCtx := &gin.Context{}
			output := api.UpdateKeyAPI(ginCtx, model.Response{Key: st.key, Value: st.keyValue[st.key]}, st.keyValue)
			assert.Equal(t, st.expectedBody, output)
			// output := api.UpdateKeyAPI(ginCtx, req, data)
			// assert.Equal(t, st.expectedBody, output)
		})
	}
}

// func TestDisplayKeys(t *testing.T) {
// 	testCases := []struct {
// 		name         string
// 		expectedCode int
// 		expectedBody map[string]string
// 	}{
// 		{
// 			name:         "display keys",
// 			expectedCode: 200,
// 			expectedBody: map[string]string{
// 				"talal": "56",
// 			},
// 		},
// 	}
// 	for _, ts := range testCases {
// 		t.Run(ts.name, func(t *testing.T) {
// 			ginCtx := &gin.Context{}
// 			output := api.DisplayKeyAPI(ginCtx, ts.expectedBody)
// 			assert.Equal(t, ts.expectedBody, output)
// 		})
// 	}
// }

func TestDisplayKeys(t *testing.T) {
	testCases := []struct {
		name         string
		expectedCode int
		expectedBody map[string]string
	}{
		{
			name:         "display keys",
			expectedCode: 200,
			expectedBody: map[string]string{
				"name": "talal",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := api.DisplayKeyAPI(tc.expectedBody)
			assert.Equal(t, tc.expectedBody, responseDataToMap(response))
		})
	}
}

func responseDataToMap(response []model.Response) map[string]string {
	result := make(map[string]string)
	for _, res := range response {
		result[res.Key] = res.Value
	}
	return result
}

func TestStoreKey(t *testing.T) {
	testCases := []struct {
		name         string
		key          string
		keyValue     map[string]string
		expectedCode int
		expectedBody map[string]string
	}{
		{
			name:         "store map data by key",
			key:          "name",
			keyValue:     map[string]string{"name": "talal"},
			expectedCode: 200,
			expectedBody: map[string]string{
				"name": "talal"},
		},
	}
	for _, sa := range testCases {
		t.Run(sa.name, func(t *testing.T) {
			ginCtx := &gin.Context{}
			output, err := api.StoreKeyAPI(ginCtx, sa.key, sa.keyValue)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			assert.Equal(t, sa.expectedBody, output)
		})
	}
}
