package controller

import (
	"goapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreKey(c *gin.Context) {

	var reqData model.Response

	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, "json binding error")
	}
	data[reqData.Key] = reqData.Value
	c.JSON(http.StatusOK, data)
}

