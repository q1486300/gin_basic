package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", m1, index, m2)

	router.Run(":8080")
}

func m1(c *gin.Context) {
	fmt.Println("m1 ...in")
	c.Next()
	fmt.Println("m1 ...out")
}

func index(c *gin.Context) {
	fmt.Println("index ...in")
	c.JSON(http.StatusOK, gin.H{"msg": "index"})
	c.Abort() // 如果其中一個中間件調用了 c.Abort()，後續中間件將不再執行
	c.Next()
	fmt.Println("index ...out")
}

func m2(c *gin.Context) {
	fmt.Println("m2 ...in")
	c.Next()
	fmt.Println("m2 ...out")
}
