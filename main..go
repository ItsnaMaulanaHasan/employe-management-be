package main

import (
	"be-employee-management/internal/routes"
	standard "be-employee-management/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, standard.Response{
			Success: true,
			Message: "Backend is running well!",
		})
	})

	routes.SetUpRoutes(r)

	r.Run(":8000")
}
