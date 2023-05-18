package controller

import (
	"goapi/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayKeys(c *gin.Context){
	var req []model.Response
	for key,value:=range data {
		req=append(req,model.Response{key,value})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "map key-values",
		"data":    req,
	})
}