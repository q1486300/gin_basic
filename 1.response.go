package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 載入模板資料夾下所有的模板文件
	router.LoadHTMLGlob("templates/*")
	// 在 gin 中，沒有相對文件的路徑，只有相對專案的路徑
	// 網頁請求這個靜態資料夾的前綴，第二個參數是一個資料夾; 前綴不要重複
	router.StaticFS("/static", http.Dir("static/static"))
	// 配置單個文件: 網頁請求的路由，文件的路徑
	router.StaticFile("/golang.jpg", "static/golang.jpg")

	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/redirect", _redirect)
	router.Run(":8080")
}

func _string(context *gin.Context) {
	context.String(http.StatusOK, "Hello!")
}

func _json(context *gin.Context) {
	// json 響應結構體
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 忽略轉換為 json
	}
	//user := UserInfo{"名稱", 26, "123456"}
	//context.JSON(http.StatusOK, user)

	// json 響應 map
	//userMap := map[string]string{
	//	"user_name": "名稱",
	//	"age":       "26",
	//}
	//context.JSON(http.StatusOK, userMap)

	// 直接響應 json
	context.JSON(http.StatusOK, gin.H{"user_name": "名稱", "age": 26})
}

func _xml(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{
		"user":    "名稱1",
		"message": "hey",
		"status":  http.StatusOK,
		"data":    gin.H{"user": "名稱2"},
	})
}

func _yaml(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{
		"user":    "名稱1",
		"message": "hey",
		"status":  http.StatusOK,
		"data":    gin.H{"user": "名稱2"},
	})
}

func _html(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{"user_name": "名稱"})
}

func _redirect(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "/html")
}
