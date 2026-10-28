package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	log.Println("Server running at: 8080")
	r.Run(":8080")
}