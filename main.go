package main

import (
	"github.com/gin-gonic/gin"
)

type createCompareReq struct {
	User           int    `json:"user"`
	BeforeUrl      string `json:"before_url"`
	AfterUrl       string `json:"after_url"`
	OptReport      int    `json:"optReport"`
	ScreenShotSpan int    `json:"screenshotSpan"`
}

type pa struct {
	User        int    `json:"user"`
	URL         string `json:"url"`
	ErrorDetect int    `json:"errorDetect"`
	Timeout     int    `json:"timeout"`
	OptReport   int    `json:"optReport"`
}

func msgList(c *gin.Context) {
	id := c.Query("id")
	// TODO: get msg list by id
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{{"id": id}},
	})
}

type optimize struct {
	User    int    `json:"user"`
	URL     string `json:"url"`
	Timeout int    `json:"timeout"`
}

func msgHaveRead(c *gin.Context) {
	msgID := c.Query("id")
	// TODO: update msg status
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": msgID,
	})
}

func createCompare(c *gin.Context) {
	var req createCompareReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}
	// TODO: create compare
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": req,
	})
}

func getCompareList(c *gin.Context) {
	// TODO: get compare list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}

func createPa(c *gin.Context) {
	var req pa
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

func getPaList(c *gin.Context) {
	// TODO: get pa list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}

func createOptimize(c *gin.Context) {
	var req optimize
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

func getOptimizeList(c *gin.Context) {
	// TODO: get optimize list
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": []gin.H{},
	})
}

func main() {
	r := gin.Default()

	msg := r.Group("/msg")
	{
		msg.GET("/list", msgList)
		msg.PUT("/:id", msgHaveRead)
	}

	compare := r.Group("/compare")
	{
		compare.POST("", createCompare)
		compare.GET("", getCompareList)
	}

	pa := r.Group("/pa")
	{
		pa.POST("", createPa)
		pa.GET("", getPaList)
	}

	optimize := r.Group("/optimize")
	{
		optimize.POST("", createOptimize)
		optimize.GET("", getOptimizeList)
	}

	err := r.Run(":6666")
	if err != nil {
		panic(err)
	}
}
