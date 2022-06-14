package middleware

import (
	"fmt"
	"tan/src/util/log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quanhengzhuang/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const REQUEST_KEY = "Cloud-Trace-ID"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		requestID := c.Request.Header.Get(REQUEST_KEY)
		if requestID == "" {
			requestID = fmt.Sprintf("%+v", uuid.New())
		}
		requestid.Set(requestID)
		fields := []zapcore.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("traceId", requestID),
			zap.String("user-agent", c.Request.UserAgent()),
		}
		log.Info(path, fields...)
		c.Next()

		end := time.Now()
		cost := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				log.Error(e)
			}
		} else {
			// 打印返回值, 耗时
			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("traceId", requestID),
				zap.Duration("cost", cost),
			}
			log.Info(path, fields...)
		}
	}
}
