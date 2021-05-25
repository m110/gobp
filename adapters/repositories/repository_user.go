package repositories

import (
	"errors"
	"time"

	"golang-clean-architecture/app"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Age       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (ur UserRepository) FindAll() ([]app.User, error) {
	var dbUsers []*User
	err := ur.db.Find(&dbUsers).Error
	if err != nil {
		return nil, err
	}

	// Mapping database-specific storage model to the business-oriented application model.
	var appUsers []app.User
	for _, u := range dbUsers {
		appUsers = append(appUsers, app.User{
			ID:        u.ID,
			Name:      u.Name,
			Age:       u.Age,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		})
	}

	return appUsers, nil
}

func (ur UserRepository) Create(u app.User) error {
	// Mapping the business-oriented application model to the database-specific storage model.
	dbUser := &User{
		ID:        u.ID,
		Name:      u.Name,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}

	if err := ur.db.Create(dbUser).Error; !errors.Is(err, nil) {
		return err
	}

	return nil
}
