package msg

import (
	config "github.com/dkZzzz/qapm_backend/config"
	request "github.com/dkZzzz/qapm_backend/request"
	"github.com/gin-gonic/gin"
)

type Message struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// MsgList 处理获取消息列表的请求
func MsgList(c *gin.Context) {
	msgID := c.Query("id")
	url := config.BaseURL + "messageList?id=" + msgID
	resp := request.GET(url)

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}

// MsgHaveRead 处理标记消息为已读的请求
func MsgHaveRead(c *gin.Context) {
	msgID := c.Query("id")

	url := config.BaseURL + "messageRead?id=" + msgID
	resp := request.PUT(url)

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}
