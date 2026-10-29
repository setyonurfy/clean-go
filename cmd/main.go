package main

import (
	"log"
	"clean-go/internal/infrastructure/pgsql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println("Starting server...")

	pgdb, err := pgsql.Init()
	if err != nil {
		panic(err)
	}
	if pgdb != nil {
		log.Println("PGSQL database initialized successfully")
		
	}
	r := gin.Default()

	log.Println("Server running at: 8080")
	r.Run(":8080")
}