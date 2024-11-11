package main

import (
	"fmt"
	"log"
	"wallet/internals/core/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	port = 8014
)

func main() {
	r := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/create-user", controllers.HandleCreateUser)

	r.Run(fmt.Sprintf(":%d", port))
}
