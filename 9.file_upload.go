package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		fmt.Println(file.Size / 1024) // 單位是 Byte
		c.SaveUploadedFile(file, "./uploads/123.jpg")

		c.JSON(http.StatusOK, gin.H{"msg": "上傳成功"})
	})

	// 用 Create 和 Copy 的方式上傳文件
	//router.POST("/upload", func(c *gin.Context) {
	//	file, _ := c.FormFile("file")
	//	readerFile, _ := file.Open()
	//	writerFile, _ := os.Create("./uploads/456.jpg")
	//	defer writerFile.Close()
	//	n, _ := io.Copy(writerFile, readerFile)
	//	fmt.Println(n)
	//
	//	c.JSON(http.StatusOK, gin.H{"msg": "上傳成功"})
	//})

	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./uploads"+file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("成功上傳 %d 個文件", len(files))})
	})

	router.Run(":8080")
}
