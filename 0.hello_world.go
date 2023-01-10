package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// go get -u github.com/gin-gonic/gin
func main() {
	// 創建一個預設的路由
	router := gin.Default()

	// 綁定路由規則和路由函數，訪問 /index 的路由，將由對應的函數去處理
	router.GET("/index", Index)

	// 啟動監聽，gin 會把 web 服務執行在本機的 0.0.0.0:8080 端口上
	router.Run(":8080")

	// 用原生 http 服務的方式，router.Run() 本質上就是 http.ListenAndServe() 的進一步封裝
	//http.ListenAndServe(":8080", router)
}

func Index(context *gin.Context) {
	context.String(http.StatusOK, "Hello World!")
}
