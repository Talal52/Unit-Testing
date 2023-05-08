package main

import (
	"net/http"
	"practice/api"

	"github.com/gin-gonic/gin"
)

type response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var data = make(map[string]string)

func getKey(c *gin.Context) {
	input := c.Param("input")
	output := api.GetKeyAPI(c, input, data)
	c.JSON(http.StatusOK, output)
}

func deleteKey(c *gin.Context) {
	input := c.Param("input")
	output, err := api.DeleteKeyAPI(c, input, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, output)
}

func updateKey(c *gin.Context) {
	var req response
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	_, exists := data[req.Key]
	if exists {
		data[req.Key] = req.Value
		c.JSON(http.StatusOK, gin.H{
			"output": "successfully updated the entered value",
			"data":   data,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "entered key is not found",
		})
	}
}

func storeKey(c *gin.Context) {
	var requestBody response
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	data[requestBody.Key] = requestBody.Value
	c.JSON(http.StatusOK, gin.H{
		"message": "Key-value pair added successfully",
		"data":    data,
	})
}

func displayKeys(c *gin.Context) {
	var responseData []response
	for key, value := range data {
		responseData = append(responseData, response{key, value})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "map key-values",
		"data":    responseData,
	})
}
func main() {
	router := gin.Default()
	router.GET("/key/:input", getKey)
	router.DELETE("/delete/:input", deleteKey)
	router.PUT("/update/:req.key", updateKey)
	router.POST("/store", storeKey)
	router.GET("/display", displayKeys)

	router.Run(":8080")
}
