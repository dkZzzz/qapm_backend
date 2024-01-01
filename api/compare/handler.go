package compare

import (
	config "github.com/dkZzzz/qapm_backend/config"
	request "github.com/dkZzzz/qapm_backend/request"
	"github.com/gin-gonic/gin"
)

type CreateCompareReq struct {
	User           int    `form:"user" json:"user"`
	BeforeUrl      string `form:"beforeUrl" json:"beforeUrl"`
	AfterUrl       string `form:"afterUrl" json:"afterUrl"`
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
		"beforeUrl":      req.BeforeUrl,
		"afterUrl":       req.AfterUrl,
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
	user := c.Query("user")
	url := config.BaseURL + "diffList?user=" + user
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
