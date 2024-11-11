package main

import (
	"fmt"
	"log"
	"wallet/internals/core/controllers"
	"wallet/internals/core/middlewares"

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

	jwtAuth, err := middlewares.NewJWTAuth("public.key")
	if err != nil {
		log.Fatalf("Error la inicializar JWTAuth: %v", err)
	}
	r.Use(jwtAuth.Middleware())

	r.POST("/create-user", controllers.HandleCreateUser)

	r.Run(fmt.Sprintf(":%d", port))
}
