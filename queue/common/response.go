package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Success(c *gin.Context, data interface{}) {
	Failed(c, "success", 200, data)
}

func Failed(c *gin.Context, msg string, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":       msg,
		"code":      code,
		"timestamp": time.Now().Unix(),
		"payload":   data,
	})
}
