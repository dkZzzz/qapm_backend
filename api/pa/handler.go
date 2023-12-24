package pa

import "github.com/gin-gonic/gin"

type Pa struct {
	User        int    `json:"user"`
	URL         string `json:"url"`
	ErrorDetect bool    `json:"errorDetect"`
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
	// TODO: create pa
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": req,
	})
}

func GetPaList(c *gin.Context) {
	// TODO: get pa list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}
