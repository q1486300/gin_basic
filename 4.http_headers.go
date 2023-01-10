package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()

	// Header 的各種取得方式
	router.GET("/", func(c *gin.Context) {
		// 首字母大小寫不區分，單字與單字之間用 - 連接
		fmt.Println(c.GetHeader("User-Agent"))
		//fmt.Println(c.GetHeader("user-agent"))
		//fmt.Println(c.GetHeader("user-Agent"))

		// Header 是一個普通的 map[string][]string
		fmt.Println(c.Request.Header)
		// 如果是使用 Get() 方法，或者是 c.GetHeader()，那麼不區分大小寫，並且返回第一個 value
		fmt.Println(c.Request.Header.Get("User-Agent"))
		fmt.Println(c.Request.Header["User-Agent"])
		// 如果是用 map 的取值方式，請注意大小寫問題 (區分大小寫)
		fmt.Println(c.Request.Header["user-agent"])

		// 自定義的 Header，用 Get() 方法也是不區分大小寫
		fmt.Println(c.Request.Header.Get("Token"))
		fmt.Println(c.Request.Header.Get("token"))

		c.JSON(http.StatusOK, gin.H{"msg": "成功"})
	})

	// 爬蟲與用戶區別對待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		// 用正則去匹配
		// 字串的包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬蟲來了
			c.JSON(http.StatusOK, gin.H{"data": "這是響應給爬蟲的資料"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "這是響應給用戶的資料"})
	})

	// 設置 Header
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "sdfhgsdjhfgsdfjh")
		c.Header("Content-Type", "application/text; charset=urf-8")
		c.JSON(http.StatusOK, gin.H{"data": "看看 Header"})
	})

	router.Run(":8080")
}
