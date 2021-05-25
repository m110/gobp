package interactors

import (
	"golang-clean-architecture/core/models"
	"golang-clean-architecture/core/ports/presenters"
	"golang-clean-architecture/core/ports/repositories"
)

type userInteractor struct {
	UserRepository repositories.UserRepository
	UserPresenter  presenters.UserPresenter
}

type UserInteractor interface {
	Get(u []*models.User) ([]*models.User, error)
	Create(u *models.User) (*models.User, error)
}

func NewUserInteractor(r repositories.UserRepository, p presenters.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*models.User) ([]*models.User, error) {
	// Business goes here

	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Create(u *models.User) (*models.User, error) {
	u, err := us.UserRepository.Create(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
