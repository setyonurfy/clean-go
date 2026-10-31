package main

import (
	"log"
	"os"

	"clean-go/internal/delivery/http/handler"
	"clean-go/internal/infrastructure/pgsql"
	"clean-go/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println("Starting server...")

	pgdb, err := pgsql.Init()
	if err != nil {
		panic(err)
	}
	log.Println("PGSQL database initialized successfully")

	r := gin.Default()

	userRepo := pgsql.NewUserRepoPG(pgdb)
	authUC := usecase.NewAuthUsecase(userRepo, os.Getenv("JWT_SECRET"))
	handler.NewAuthHandler(r, authUC)

	log.Println("Server running at: 8080")
	r.Run(":8080")
}
