package main

import (
	"github.com/dkZzzz/qapm_backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRoutes(r)
	err := r.Run(":6666")
	if err != nil {
		panic(err)
	}
}
