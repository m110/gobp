package app

import (
	"time"
)

// User is the application model. It should not be influenced in any way by database or HTTP implementation details.
type User struct {
	ID        uint
	Name      string
	Age       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Previously in ports. But there's no reason to put the interface in a separate layer.
// Because of Go's implicit interfaces you can keep them close to what uses them.
type userRepository interface {
	FindAll() ([]User, error)
	Create(u User) error
}

type UserInteractor struct {
	repository userRepository
}

func NewUserInteractor(repository userRepository) UserInteractor {
	return UserInteractor{
		repository,
	}
}

func (us UserInteractor) Get() ([]User, error) {
	users, err := us.repository.FindAll()
	if err != nil {
		return nil, err
	}

	// Previously in presenters
	for _, u := range users {
		u.Name = "Mr." + u.Name
	}

	return users, nil
}

func (us UserInteractor) Create(u User) (User, error) {
	err := us.repository.Create(u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
