package pa

import (
	"github.com/dkZzzz/qapm_backend/config"
	"github.com/dkZzzz/qapm_backend/request"
	"github.com/gin-gonic/gin"
)

type Pa struct {
	User        int    `json:"user"`
	URL         string `json:"url"`
	ErrorDetect int    `json:"errorDetect"`
	Timeout     int    `json:"timeout"`
	OptReport   int    `json:"optReport"`
}

func CreatePa(c *gin.Context) {
	var req Pa
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}

	url := config.BaseURL + "performance"
	BodyData := map[string]interface{}{
		"user":        req.User,
		"url":         req.URL,
		"errorDetect": req.ErrorDetect,
		"timeout":     req.Timeout,
		"optReport":   req.OptReport,
	}
	resp := request.POST(url, BodyData)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}

func GetPaList(c *gin.Context) {
	url := config.BaseURL + "performanceList"
	resp := request.GET(url)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}

func GetPaDetail(c *gin.Context) {
	id := c.Param("id")
	url := config.BaseURL + "performanceDetail?id=" + id
	resp := request.GET(url)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}
