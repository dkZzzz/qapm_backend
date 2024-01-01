package compare

import (
	config "github.com/dkZzzz/qapm_backend/config"
	request "github.com/dkZzzz/qapm_backend/request"
	"github.com/gin-gonic/gin"
)

type CreateCompareReq struct {
	User           int    `form:"user" json:"user"`
	BeforeUrl      string `form:"before_url" json:"before_url"`
	AfterUrl       string `form:"after_url" json:"after_url"`
	OptReport      bool   `form:"optReport" json:"optReport"`
	ScreenShotSpan int    `form:"screenshotSpan" json:"screenshotSpan"`
}

func CreateCompare(c *gin.Context) {
	var req CreateCompareReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}

	BodyData := map[string]interface{}{
		"user":           req.User,
		"before_url":     req.BeforeUrl,
		"after_url":      req.AfterUrl,
		"optReport":      req.OptReport,
		"screenshotSpan": req.ScreenShotSpan,
	}

	url := config.BaseURL + "diff"

	resp := request.POST(url, BodyData)

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}

func GetCompareList(c *gin.Context) {
	url := config.BaseURL + "diffList"
	resp := request.GET(url)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}

func GetCompareDetail(c *gin.Context) {
	id := c.Param("id")
	url := config.BaseURL + "diffDetail?id=" + id
	resp := request.GET(url)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": resp,
	})
}
