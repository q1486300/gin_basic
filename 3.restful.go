package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/articles", _getList)
	router.GET("/articles/:id", _getDetail)
	router.POST("/articles", _create)
	router.PUT("/articles/:id", _update)
	router.DELETE("/articles/:id", _delete)

	router.Run(":8080")
}

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _getList(c *gin.Context) {
	// 包含搜尋、分頁
	articleList := []ArticleModel{
		{"Go語言入門", "這篇文章是Go語言入門"},
		{"Gin入門", "這篇文章是Gin入門"},
		{"Gorm入門", "這篇文章是Gorm入門"},
	}
	c.JSON(http.StatusOK, Response{0, articleList, "成功"})
}

func _getDetail(c *gin.Context) {
	// 取得 param 中的 id
	fmt.Println(c.Param("id"))
	article := ArticleModel{"Go語言入門", "這篇文章是Go語言入門"}
	c.JSON(http.StatusOK, Response{0, article, "成功"})
}

func _create(c *gin.Context) {
	// 取得前端傳來的 json 資料
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, Response{0, article, "添加成功"})
}

func _update(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, Response{0, article, "修改成功"})
}

func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(http.StatusOK, Response{0, gin.H{}, "刪除成功"})
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			return err
		}
	}
	return nil
}
