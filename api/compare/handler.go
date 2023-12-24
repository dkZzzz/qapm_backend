package compare

import (
	"github.com/gin-gonic/gin"
)

type CreateCompareReq struct {
	User           int    `json:"user"`
	BeforeUrl      string `json:"before_url"`
	AfterUrl       string `json:"after_url"`
	OptReport      bool    `json:"optReport"`
	ScreenShotSpan int    `json:"screenshotSpan"`
}

func CreateCompare(c *gin.Context) {
	var req CreateCompareReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": req,
	})
}

func GetCompareList(c *gin.Context) {
	// TODO: get compare list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}
