package repository

import (
	"golang_wallet/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByIDForUpdate(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}
