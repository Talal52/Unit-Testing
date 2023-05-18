package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteKey(c *gin.Context){
	key := c.Param("key")
	_,ok:=data[key]
	if !ok{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "entered key not found",
		})
	}
	DeleteKeyAPI(key)
	delete(data, key)
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
		"data":    data,
	})
}
func DeleteKeyAPI(input string) map[string]string{
	return data
}