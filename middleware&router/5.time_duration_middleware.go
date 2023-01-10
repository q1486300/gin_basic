package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	timeDuration := api.Group("").Use(TimeMiddleware)
	{
		timeDuration.GET("/test", func(c *gin.Context) {
			time.Sleep(1 * time.Second)
			c.JSON(http.StatusOK, gin.H{"msg": "請求成功"})
		})
	}

	router.Run(":8080")
}

func TimeMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	since := time.Since(startTime)
	// 取得目前請求對應的函數
	f := c.HandlerName()
	fmt.Printf("函數 %s 耗時 %v\n", f, since)
}
