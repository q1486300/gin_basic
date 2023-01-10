package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func LogMiddleware(c *gin.Context) {
	startTime := time.Now()

	c.Next()

	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)

	reqMethod := formatMethod(c.Request.Method)

	reqUri := c.Request.RequestURI

	statusCode := formatStatusCode(c.Writer.Status())

	clientIP := c.ClientIP()

	logrus.WithField("type", "builtin").Infof("[GIN] %s |%s| %13v | %15s | %s \"%s\" ",
		startTime.Format("2006-01-02 - 15:04:05"),
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri,
	)
}

func formatMethod(method string) string {
	methodColor := reset

	switch method {
	case http.MethodGet:
		methodColor = blue
	case http.MethodPost:
		methodColor = cyan
	case http.MethodPut:
		methodColor = yellow
	case http.MethodDelete:
		methodColor = red
	case http.MethodPatch:
		methodColor = green
	case http.MethodHead:
		methodColor = magenta
	case http.MethodOptions:
		methodColor = white
	}

	return fmt.Sprintf("%s %s    %s", methodColor, method, reset)
}

func formatStatusCode(statusCode int) string {
	statusCodeColor := reset

	switch {
	case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
		statusCodeColor = green
	case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
		statusCodeColor = white
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		statusCodeColor = yellow
	}

	return fmt.Sprintf("%s %d %s", statusCodeColor, statusCode, reset)
}
