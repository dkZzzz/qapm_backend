package optimize

import (
	"strconv"

	"github.com/dkZzzz/qapm_backend/models"
	"github.com/gin-gonic/gin"
)

type Optimize struct {
	User       int    `form:"user" json:"user"`
	TaskId     string `form:"taskId" json:"taskId"`
	GptVersion string `form:"gptVersion" json:"gptVersion"`
	Url        string `form:"url" json:"url"`
}

func CreateOptimize(c *gin.Context) {
	var req Optimize
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": err.Error(),
		})
		return
	}

	id, err := models.CreateOptimize(&models.Optimize{
		User:       req.User,
		TaskId:     req.TaskId,
		GptVersion: req.GptVersion,
		Url:        req.Url})

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
		"msg":  "创建成功",
		"data": gin.H{
			"id": id,
		},
	})
}

func GetOptimizeByUser(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": "user is required",
		})
		return
	}
	userID, err := strconv.Atoi(user)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": "invalid user ID",
		})
		return
	}
	optimizes, err := models.SelectOptimizeByUser(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "server error",
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{"optimizes": optimizes},
	})
}

func GetOptimizeById(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 1,
			"msg":  "bad request",
			"data": "invalid optimize ID",
		})
		return
	}
	optimizes, err := models.SelectOptimizeById(ID)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 1,
			"msg":  "server error",
			"data": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{"optimizes": optimizes},
	})
}
