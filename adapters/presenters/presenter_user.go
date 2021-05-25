package presenters

import (
	"golang-clean-architecture/core/models"
	"golang-clean-architecture/core/ports/presenters"
)

type userPresenter struct{}

func NewUserPresenter() presenters.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*models.User) []*models.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
