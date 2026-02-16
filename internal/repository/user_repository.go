package repository

import "golang_wallet/internal/domain"

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	FindByIDForUpdate(id uint) (*domain.User, error)
	Update(user *domain.User) error
}
