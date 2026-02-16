package handler

import (
	"net/http"
	"strconv"

	"golang_wallet/internal/usecase"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	uc *usecase.WalletUsecase
}

func NewWalletHandler(r *gin.Engine, uc *usecase.WalletUsecase) {
	handler := &WalletHandler{uc}

	r.POST("/withdraw", handler.Withdraw)
	r.GET("/balance/:id", handler.GetBalance)
}

type WithdrawRequest struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *WalletHandler) Withdraw(c *gin.Context) {
	var req WithdrawRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.uc.Withdraw(req.UserID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.uc.GetBalance(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"balance": user.Balance,
	})
}
