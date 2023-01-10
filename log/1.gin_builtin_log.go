package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(LogFormatterParams))
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: LogFormatterParams}))

	router.GET("/index", func(c *gin.Context) {})

	for _, info := range router.Routes() {
		fmt.Println(info.Path, info.Method, info.Handler)
	}

	router.Run(":8080")
}

func LogFormatterParams(params gin.LogFormatterParams) string {
	return fmt.Sprintf("[Custom] %s\t|%s %d %s|\t%s %s %s\t%s\n",
		params.TimeStamp.Format("2006-01-02 15:04:05"),
		params.StatusCodeColor(), params.StatusCode, params.ResetColor(),
		params.MethodColor(), params.Method, params.ResetColor(),
		params.Path,
	)
}
