package repositories

import (
	"errors"
	"golang-clean-architecture/core/models"
	"golang-clean-architecture/core/ports/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*models.User) ([]*models.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *models.User) (*models.User, error) {
	if err := ur.db.Create(u).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}
