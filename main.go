package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"

	"Fokkely-core/core"
)

var k map[string]map[string]interface{} //准备换为DataMap
var skillDict core.SkillMap

func handleHello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"each": "hello",
	})
}

// get-skill
func handleGetSkill(c *gin.Context) {
	var jsonParams map[string]string

	err := c.BindJSON(&jsonParams)
	fmt.Println(jsonParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, ok := jsonParams["id"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的字段"})
		return
	}
	sk, _ := skillDict.Get(id)
	c.JSON(http.StatusOK, sk)
}
func main() {

	app := gin.Default()
	skillDict.Load()

	/*

		test pass
		var t core.SkillMap
		t.Load()
		fmt.Println(t.Get("sk-1"))

	*/

	app.POST("/hello", handleHello)
	app.POST("/get/skill", handleGetSkill)

	// 启动服务器，默认在 :8080 上监听
	app.Run(":5700")
}
