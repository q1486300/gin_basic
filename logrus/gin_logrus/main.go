package main

import (
	"gin_basic/logrus/gin_logrus/log"
	"gin_basic/logrus/gin_logrus/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.InitLog("logrus/gin_logrus/logs", "server")

	router := gin.New()

	router.Use(middleware.LogMiddleware)

	router.GET("/", func(c *gin.Context) {
		logrus.Info("info")
		c.JSON(http.StatusOK, gin.H{"msg": "Hello"})
	})

	router.Run(":8080")
}
