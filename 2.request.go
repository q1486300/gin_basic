package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":8080")
}

// 查詢參數
// http://localhost:8080/query?
func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))   // 第二個返回值代表是否有 user 這個 query 參數
	fmt.Println(c.QueryArray("user")) // 拿到多個相同的查詢參數，?user=名稱1&user=名稱2
	fmt.Println(c.QueryMap("user"))   // ?user[id]=2&user[user]=名稱
	fmt.Println(c.DefaultQuery("user", "名稱"))
}

// 動態參數
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表單參數
// multipart/form-data 或 application/x-www-form-urlencoded
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("name", "名稱"))
	forms, err := c.MultipartForm() // 接收所有的 form 參數，包括文件
	fmt.Println(forms, err)
}

// 原始參數
func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Addr string `json:"addr"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}

func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	}
	return nil
}
