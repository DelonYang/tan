package ctrl

import (
	"net/http"
	"tan/src/model"

	"github.com/gin-gonic/gin"
)

const (
	RET_CODE = "retCode"
	RET_MSG  = "retMsg"
	RET_DATA = "retData"
)

func WriteJson(c *gin.Context, h gin.H) {
	if _, ok := h[RET_CODE]; !ok {
		h[RET_CODE] = model.ErrCodeNo
	}

	if _, ok := h[RET_MSG]; !ok {
		if s, ok := model.ErrCodeName[h[RET_CODE].(model.ErrCode)]; ok {
			h[RET_MSG] = s
		} else {
			h[RET_MSG] = "unknown error"
		}
	}

	c.JSON(http.StatusOK, h)
}
