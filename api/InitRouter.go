package api

import (
	"wallet/internal/core/handlers"
	"wallet/internal/core/services"
	repository "wallet/internal/repository/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewGormRepository(db)
	userService := services.NewUserService(userRepo)
	handler := handlers.NewWalletHandler(userService)

	r.POST("/create-user", handler.CreateUser)
	r.POST("/create-transaction", handler.CreateTransaction)
	r.POST("/generate-excel", handler.GenerateExcel)
}
