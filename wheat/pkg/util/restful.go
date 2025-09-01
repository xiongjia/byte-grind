package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() (http.Handler, error) {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	return router.Handler(), nil
}
