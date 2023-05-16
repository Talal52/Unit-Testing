package main

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	data := make(map[string]string)
	router := gin.Default()
	router.GET("/display/:key", controller.GetKey)
	router.DELETE("/delete/:key", controller.DeleteKey)
	router.PUT("/update", controller.UpdateKey)
	router.POST("/store", controller.StoreKey)
	//router.GET("/display", displayKeys)
	router.GET("/display", func(c *gin.Context) {
		data = map[string]string{"name": "talal"}
		controller.DisplayKeys(c, data)
	})

	router.Run(":8080")
}
