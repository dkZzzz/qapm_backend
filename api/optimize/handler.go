package optimize

import "github.com/gin-gonic/gin"

type Optimize struct {
	User    int    `json:"user"`
	URL     string `json:"url"`
	Timeout int    `json:"timeout"`
}

func CreateOptimize(c *gin.Context) {
	var req Optimize
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}
	// TODO: create optimize
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": req,
	})
}

func GetOptimizeList(c *gin.Context) {
	// TODO: get optimize list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}
