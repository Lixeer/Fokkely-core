package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"

	"Fokkely-core/core"
)

func processData(c *gin.Context) {
	// 声明一个结构体来映射JSON数据
	var data map[string]string
	fmt.Println(data)
	// 解析请求的JSON数据到data结构体
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从data中获取"each"字段的值
	eachValue, ok := data["each"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'each' field in JSON data"})
		return
	}

	// 输出数据到响应
	c.JSON(http.StatusOK, gin.H{"message": "Received POST request", "eachValue": eachValue})
}

func handleHello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"each": "hello",
	})
}

// get-skill
func handleGetSkill(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"each": "hello",
	})
}
func main() {

	app := gin.Default()
	k, _ := core.GetSkillData()
	fmt.Println("k:", k["sk-1"])

	app.POST("/hello", handleHello)
	app.POST("/process", processData)
	app.POST("/get-skill", handleGetSkill)

	// 启动服务器，默认在 :8080 上监听
	app.Run(":5700")
}
