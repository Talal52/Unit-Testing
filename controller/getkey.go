package controller

import (
	"goapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKey(c *gin.Context) {
	var reqData model.Response

	key := c.Param("key")
	resp, ok := data[key]
	if !ok {
		c.JSON(http.StatusNotFound, "No key found")
		return
	}

	GettingKeyAPI(key)
	data[key] = reqData.Value
	c.JSON(http.StatusOK, resp)
}

func GettingKeyAPI(input string) string {
	return data["input"]
}
