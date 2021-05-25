package repositories

import (
	"golang-clean-architecture/core/models"
)

type UserRepository interface {
	FindAll(u []*models.User) ([]*models.User, error)
	Create(u *models.User) (*models.User, error)
}
