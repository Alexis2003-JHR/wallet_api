package main

import (
	"fmt"
	"log"
	database "wallet/cmd/database"
	handlers "wallet/internal/core/handlers"
	"wallet/internal/core/middlewares"
	"wallet/internal/core/services"
	repository "wallet/internal/repository/user"

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
	userRepo := repository.NewGormRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/create-user", userHandler.HandleCreateUser)

	r.Run(fmt.Sprintf(":%d", port))
}
