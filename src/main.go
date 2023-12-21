
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	
	app := gin.Default()

	
	app.POST("/hello", func(c *gin.Context) {
	
		c.JSON(http.StatusOK, gin.H{
			"each": "hello",
		})
	})

	// 启动服务器，默认在 :8080 上监听
	app.Run()
}