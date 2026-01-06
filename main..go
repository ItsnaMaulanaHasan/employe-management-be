package main

import (
	"be-employee-management/internal/database"
	"be-employee-management/internal/routes"
	standard "be-employee-management/pkg/response"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := database.ConnectPostgres(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, standard.Response{
			Success: true,
			Message: "Backend is running well!",
		})
	})

	routes.SetUpRoutes(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
