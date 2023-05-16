package controller

import (
	"goapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = make(map[string]string)

func UpdateKey(c *gin.Context) {
	var req model.Response
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input 1",
		})
		return
	}
	_, exists := data[req.Key]
	if exists {
		data[req.Key] = req.Value
		updatedValue := req.Value
		data[req.Key] = updatedValue
		c.JSON(http.StatusOK, gin.H{
			"output": "successfully update the entered value",
			"data":   data,
		})
	}
}
