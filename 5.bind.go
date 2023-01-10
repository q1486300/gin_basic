package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 綁定 json 參數，結構體中對應的 tag 為 json
	router.POST("/", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"msg": "出錯了"})
			return
		}
		c.JSON(http.StatusOK, userInfo)
	})

	// 綁定 query 參數，結構體中對應的 tag 為 form
	router.POST("/query", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"msg": "出錯了"})
			return
		}
		c.JSON(http.StatusOK, userInfo)
	})

	// 綁定 uri 參數，結構體中對應的 tag 為 uri
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"msg": "出錯了"})
			return
		}
		c.JSON(http.StatusOK, userInfo)
	})

	// 綁定 form (multipart/form-data 或 application/x-www-form-urlencoded) 參數，結構體中對應的 tag 為 form
	router.POST("/form", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"msg": "出錯了"})
			return
		}
		c.JSON(http.StatusOK, userInfo)
	})

	router.Run(":8080")
}

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}
