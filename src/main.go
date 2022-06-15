package main

import (
	"fmt"
	"tan/src/config"
	"tan/src/router"
	"tan/src/util/log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger := log.InitLog()
	defer logger.Sync()

	r := router.NewRouter()
	r.Run(fmt.Sprintf(":%d", config.Conf.Common.Port))
}
