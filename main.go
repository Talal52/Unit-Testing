package main

import (
	"fmt"
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
	router.GET("/display", controller.DisplayKeys)
	fmt.Println(data)
	router.Run(":8080")
}
