package controller

import (
	"goapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = make(map[string]string)

func UpdateKey(c *gin.Context) {
	var ver model.Vertex
	if err := c.BindJSON(&ver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	_, exists := data[ver.Key]
	if exists {
		// data[req.Key] = req.Value
		// updatedValue := req.Value
		// data[req.Key] = updatedValue
		c.JSON(http.StatusOK, gin.H{
			"output": "successfully update the entered value",
			"data":   data,
		})
	}
	UpdatingKeyAPI(ver.Key,ver.Value,ver.UpdatedValue)
	data[ver.Key]=ver.Value
	ver.UpdatedValue=ver.Value
	data[ver.Key]=ver.UpdatedValue
}
func UpdatingKeyAPI(key,value,updatedValue string) map[string]string{
	// data[key] = value
	// updatedValue = value
	data[key] = updatedValue
	return data

}