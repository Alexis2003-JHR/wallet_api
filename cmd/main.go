package main

import (
	"fmt"
	"log"
	"wallet/api"
	database "wallet/cmd/database"
	"wallet/internal/core/middlewares"

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

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("Error al inicializar base de datos %v", err)
	}

	api.InitRouter(r, db)

	r.Run(fmt.Sprintf(":%d", port))
}
