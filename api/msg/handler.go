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
	// TODO: 实际业务逻辑，从数据库或其他地方获取消息列表
	messages := []Message{
		{ID: "1", Title: "Welcome", Body: "Hello, welcome to our system!"},
		{ID: "2", Title: "New Feature", Body: "Check out the latest features."},
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": messages,
	})
}

// MsgHaveRead 处理标记消息为已读的请求
func MsgHaveRead(c *gin.Context) {
	msgID := c.Param("id")

	// TODO: 实际业务逻辑，更新数据库或其他地方的消息状态

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": msgID,
	})
}
