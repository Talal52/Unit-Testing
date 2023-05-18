package api

import (
	"goapi/model"

	"github.com/gin-gonic/gin"
)

func GetKeyAPI(c *gin.Context, input string, data map[string]string) string {
	value, exists := data[input]
	if exists {
		return value

	} else {

		return "entered Key is not found!!!"
	}
}

// func DisplayKeyAPI(c *gin.Context, data map[string]string) map[string]string {
// 	var responseData []model.Response
// 	return data
// 	for key, value := range data {
// 		responseData = append(responseData, model.Response{key, value})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "map key-values",
// 		"data":    responseData,
// 	})
// }

// func DisplayKeyAPI(c *gin.Context, data map[string]string) {
// 	var responseData []model.Response
// 	for key, value := range data {
// 		responseData = append(responseData, model.Response{key, value})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "map key-values",
// 		"data":    responseData,
// 	})
// }

// func DisplayKeyAPI(c *gin.Context, data map[string]string) []model.Response {
// 	var responseData []model.Response
// 	for key, value := range data {
// 		responseData = append(responseData, model.Response{key, value})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "map key-values",
// 		"data":    responseData,
// 	})

// 	return responseData
// }

func DisplayKeyAPI(data map[string]string) []model.Response {
	var responseData []model.Response
	for key, value := range data {
		responseData = append(responseData, model.Response{key, value})
	}
	return responseData
}

func DeleteKeyAPI(c *gin.Context, input string, data map[string]string) (interface{}, error) {
	_, exists := data[input]
	if exists {
		delete(data, input)
		return data, nil

	} else {
		
		return "Entered key not found", nil
	}
}
func UpdateKeyAPI(c *gin.Context, req model.Response, data map[string]string) map[string]string {
	_, exists := data[req.Key]
	if exists {
		data[req.Key] = req.Value
		return data
	} else {
		return map[string]string{"error": "key not found"}
	}
}

func StoreKeyAPI(c *gin.Context, input string, data map[string]string) (map[string]string, error) {
	_, exists := data[input]
	if exists {
		return data, nil
	}
	return data, nil
}
