package queue

import (
	"github.com/gin-gonic/gin"
	"go-toolbox/queue/handler"
	"go-toolbox/queue/node"
)

func Init(gin *gin.Engine) *node.Queue {
	q := node.NewQueue()
	handler.Queue = q
	if gin != nil {
		InitRouter(gin)
	}
	go q.PutQueueChan()
	return q
}

func InitRouter(gin *gin.Engine) {
	group := gin.Group("queue")
	initQueueRouter(group)
}
func initQueueRouter(group *gin.RouterGroup) {
	group.POST("put", handler.Put)
}
