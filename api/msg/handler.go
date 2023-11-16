package msg

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// MsgList 处理获取消息列表的请求
func MsgList(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{},
	})
}

// MsgHaveRead 处理标记消息为已读的请求
func MsgHaveRead(c *gin.Context) {
	msgID := c.Param("id")

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{"msgID": msgID},
	})
}
