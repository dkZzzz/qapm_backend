package router

import (
	"github.com/dkZzzz/qapm_backend/api/compare"
	"github.com/dkZzzz/qapm_backend/api/msg"
	"github.com/dkZzzz/qapm_backend/api/optimize"
	"github.com/dkZzzz/qapm_backend/api/pa"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	msgGroup := api.Group("/msg")
	{
		msgGroup.GET("/list", msg.MsgList)
		msgGroup.PUT("", msg.MsgHaveRead)
	}

	compareGroup := api.Group("/compare")
	{
		compareGroup.POST("", compare.CreateCompare)
		compareGroup.GET("", compare.GetCompareList)
		compareGroup.GET("/:id", compare.GetCompareDetail)
	}

	paGroup := api.Group("/pa")
	{
		paGroup.POST("", pa.CreatePa)
		paGroup.GET("", pa.GetPaList)
		paGroup.GET("/:id", pa.GetPaDetail)
	}

	optimizeGroup := api.Group("/optimize")
	{
		optimizeGroup.POST("", optimize.CreateOptimize)
		optimizeGroup.GET("", optimize.GetOptimizeByUser)
		optimizeGroup.GET("/:id", optimize.GetOptimizeById)
	}
}
