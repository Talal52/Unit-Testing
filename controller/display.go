package controller

import (
	"goapi/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayKeys(c *gin.Context, data map[string]string) {
	responseData := api.DisplayKeyAPI(data)
	c.JSON(http.StatusOK, responseData)
}
