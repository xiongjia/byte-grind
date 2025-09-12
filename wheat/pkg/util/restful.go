package util

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		reqUrl := c.Request.RequestURI
		slog.Debug("custom middleware", slog.Duration("latancy", latency), slog.String("req", reqUrl))
	}
}

func NewGinRouter() (http.Handler, error) {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(customMiddleware())
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	return router.Handler(), nil
}
