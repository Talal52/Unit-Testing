package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKey(c *gin.Context) {
	key := c.Param("key")
	resp, ok := data[key]
	if !ok {
		c.JSON(http.StatusNotFound, "No key found")
	}

	c.JSON(http.StatusOK, resp)
}

