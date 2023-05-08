package api

import "github.com/gin-gonic/gin"

func GetKeyAPI(c *gin.Context, input string, data map[string]string) string {
	value, exists := data[input]
	if exists {
		return value
		// c.JSON(http.StatusOK, gin.H{
		// 	"key":   input,
		// 	"value": value,
		// })
	} else {
		// c.JSON(http.StatusNotFound, gin.H{
		// "error message": "entered Key is not found!!!",
		// })
		return "entered Key is not found!!!"
	}
}

func DeleteKeyAPI(c *gin.Context, input string, data map[string]string) (interface{}, error) {
	_, exists := data[input]
	if exists {
		delete(data, input)
		return data, nil
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "deleted successfully",
		// 	"data":    data,
		// })
	} else {
		// c.JSON(http.StatusNotFound, gin.H{
		// 	"error": "entered key not found",
		// })
		return "Entered key not found", nil
	}
}
