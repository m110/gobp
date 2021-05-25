package presenters

import (
	"golang-clean-architecture/core/models"
)

type UserPresenter interface {
	ResponseUsers(u []*models.User) []*models.User
}
