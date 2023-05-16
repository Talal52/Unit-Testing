package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteKey(c *gin.Context){
	key := c.Param("key")
		_, exists := data[key]
		if exists {
			delete(data, key)
			c.JSON(http.StatusOK, gin.H{
				"message": "deleted successfully",
				"data":    data,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "entered key not found",
			})
		}
}