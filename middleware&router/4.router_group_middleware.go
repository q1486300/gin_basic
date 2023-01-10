package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "1234"})
	})

	_UserRouterInit(api)

	router.Run(":8080")
}

func _UserRouterInit(router *gin.RouterGroup) {
	//userManage := router.Group("/user_manage").Use(Middleware)
	userManage := router.Group("/user_manage").Use(Middleware("用戶驗證失敗"))
	{
		userManage.GET("/users", _UserListView)
	}
}

//func Middleware(c *gin.Context) {
//	token := c.GetHeader("token")
//	if token == "1234" {
//		c.Next()
//		return
//	}
//	c.JSON(http.StatusOK, Res{1001, nil, "權限驗證失敗"})
//	c.Abort()
//}

// Middleware 傳入參數自定義錯誤訊息
func Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(http.StatusOK, Res{1001, nil, msg})
		c.Abort()
	}
}

type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _UserListView(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	userList := []UserInfo{
		{"名稱1", 26},
		{"名稱2", 23},
		{"名稱3", 20},
	}
	c.JSON(http.StatusOK, Res{0, userList, "請求成功"})
}
