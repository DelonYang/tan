package ctrl

import (
	"tan/src/util/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ping(c *gin.Context) {
	var res = gin.H{}
	defer WriteJson(c, res)
	log.Info("info test", zap.String("ccc", "aksd"))
	log.Warn("info test", zap.String("ccc", "aksd"))
	log.Error("info test", zap.String("ccc", "aksd"))
}
