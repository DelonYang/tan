package router

import (
	"tan/src/ctrl"
	"tan/src/middleware"

	"github.com/gin-gonic/gin"
)

const REQUEST_KEY = "Cloud-Trace-ID"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recover())

	r.GET("/ping", ctrl.Ping)
	r.GET("/panic", func(c *gin.Context) {
		panic("asdljfals")
	})
	return r
}
