package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	UserRouterInit(api)
	ArticleRouterInit(api)

	router.Run(":8080")
}

func UserRouterInit(router *gin.RouterGroup) {
	userManage := router.Group("/user_manage")
	{
		userManage.GET("/users", UserListView)
	}
}

func ArticleRouterInit(router *gin.RouterGroup) {
	articleManage := router.Group("/article_manage")
	{
		articleManage.GET("/articles", ArticleListView)
	}
}

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	userList := []UserInfo{
		{"名稱1", 26},
		{"名稱2", 23},
		{"名稱3", 20},
	}
	c.JSON(http.StatusOK, Response{0, userList, "請求成功"})
}

func ArticleListView(c *gin.Context) {
	articleList := []ArticleInfo{
		{"Go語言入門", "這篇文章是Go語言入門"},
		{"Gin入門", "這篇文章是Gin入門"},
		{"Gorm入門", "這篇文章是Gorm入門"},
	}
	c.JSON(http.StatusOK, Response{0, articleList, "請求成功"})
}
