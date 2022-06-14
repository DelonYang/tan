package router

import (
	"net/http"

	"tan/src/middleware"
	"tan/src/util/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const REQUEST_KEY = "Cloud-Trace-ID"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recover())

	r.GET("/ping", func(c *gin.Context) {
		log.Info("info test", zap.String("ccc", "aksd"))
		log.Warn("info test", zap.String("ccc", "aksd"))
		log.Error("info test", zap.String("ccc", "aksd"))
		c.String(http.StatusOK, "pong")
	})
	r.GET("/panic", func(c *gin.Context) {
		panic("asdljfals")
	})
	return r
}
