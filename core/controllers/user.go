package controllers

import (
	"errors"
	"net/http"

	"golang-clean-architecture/core/interactors"
	"golang-clean-architecture/core/models"
)

type UserController interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
}

type userController struct {
	userInteractor interactors.UserInteractor
}

func NewUserController(us interactors.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*models.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(c Context) error {
	var params models.User

	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	u, err := uc.userInteractor.Create(&params)
	if !errors.Is(err, nil) {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
