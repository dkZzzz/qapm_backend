package main

import (
	"os"

	"github.com/dkZzzz/qapm_backend/internal/router"
	"github.com/dkZzzz/qapm_backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer models.DB.Close()

	gin.DefaultWriter = os.Stdout
	r := gin.Default()
	router.RegisterRoutes(r)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
