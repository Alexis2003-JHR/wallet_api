package handlers

import "wallet/internal/core/services"

type WalletHandler struct {
	service *services.UserService
}

func NewWalletHandler(service *services.UserService) *WalletHandler {
	return &WalletHandler{service: service}
}
