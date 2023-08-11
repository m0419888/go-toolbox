package go_toolbox

import (
	"github.com/gin-gonic/gin"
	"go-toolbox/queue"
	"go-toolbox/queue/node"
)

func NewQueue(gin *gin.Engine) *node.Queue {
	return queue.Init(gin)
}
