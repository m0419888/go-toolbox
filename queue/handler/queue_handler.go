package handler

import (
	"github.com/gin-gonic/gin"
	"go-toolbox/queue/common"
	"go-toolbox/queue/node"
	"net/http"
)

var Queue *node.Queue

func Put(context *gin.Context) {
	var body common.PutBody
	err := context.BindJSON(&body)
	if err != nil {
		common.Failed(context, "参数错误", http.StatusBadRequest, nil)
	}
	Queue.PutChan(body)
	common.Success(context, "ok")
}
