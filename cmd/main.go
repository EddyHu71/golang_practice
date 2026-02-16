package main

import (
	"golang_wallet/internal/handler"
	"golang_wallet/internal/repository"
	"golang_wallet/internal/usecase"
	"golang_wallet/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	db := pkg.InitDB()

	userRepo := repository.NewUserRepository(db)
	walletUC := usecase.NewWalletUsecase(userRepo, db)

	r := gin.Default()

	handler.NewWalletHandler(r, walletUC)

	r.Run(":8080")
}
