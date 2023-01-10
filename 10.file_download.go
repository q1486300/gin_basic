package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 注意: 文件下載瀏覽器可能會有緩存，解決辦法是在 URL 後面加查詢參數
	router.POST("/download1", func(c *gin.Context) {
		// 表示是文件流，喚起瀏覽器下載，一般設置了這個，就要設置文件名
		c.Header("Content-Type", "application/octet-stream")
		// 用來指定下載下來的文件名
		c.Header("Content-Disposition", "attachment; filename="+"這裡填入文件名.jpg")
		// 表示傳輸過程中的編碼形式，避免有亂碼問題
		c.Header("Content-Transfer-Encoding", "binary")

		c.File("./uploads/golang.jpg")
	})

	// 前後端模式下的文件下載
	router.POST("/download2", func(c *gin.Context) {
		c.Header("fileName", "xxx.jpg")
		c.Header("msg", "文件下載成功")

		c.File("./uploads/golang.jpg")
	})

	router.Run(":8080")
}
