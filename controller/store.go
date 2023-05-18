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

	StoringDataAPI(reqData.Key, reqData.Value)
	data[reqData.Key] = reqData.Value
	c.JSON(http.StatusOK, data)
}

func StoringDataAPI(key, value string) map[string]string {
	data[key] = value
	return data
}
