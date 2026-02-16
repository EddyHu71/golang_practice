package usecase

import (
	"errors"
	"golang_wallet/internal/domain"
	"golang_wallet/internal/repository"

	"gorm.io/gorm"
)

type WalletUsecase struct {
	userRepo repository.UserRepository
	db       *gorm.DB
}

func NewWalletUsecase(repo repository.UserRepository, db *gorm.DB) *WalletUsecase {
	return &WalletUsecase{
		userRepo: repo,
		db:       db,
	}
}

func (u *WalletUsecase) Withdraw(userID uint, amount float64) error {
	return u.db.Transaction(func(tx *gorm.DB) error {

		user, err := u.userRepo.FindByIDForUpdate(userID)
		if err != nil {
			return err
		}

		if user.Balance < amount {
			return errors.New("insufficient balance")
		}

		user.Balance -= amount

		return u.userRepo.Update(user)
	})
}

func (u *WalletUsecase) GetBalance(userID uint) (*domain.User, error) {
	return u.userRepo.FindByID(userID)
}
