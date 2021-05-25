package http

import (
	"errors"
	"net/http"
	"time"

	"golang-clean-architecture/app"

	"github.com/labstack/echo"
)

type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       string     `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserController struct {
	userInteractor app.UserInteractor
}

func NewUserController(us app.UserInteractor) UserController {
	return UserController{us}
}

// No problem with depending on echo.Context here - it's the http layer
func (uc UserController) GetUsers(c echo.Context) error {
	appUsers, err := uc.userInteractor.Get()
	if err != nil {
		return err
	}

	var httpUsers []User
	for _, u := range appUsers {
		httpUsers = append(httpUsers, userToHTTPResponse(u))
	}

	return c.JSON(http.StatusOK, httpUsers)
}

func (uc UserController) CreateUser(c echo.Context) error {
	var params User
	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	appUser := httpRequestToUser(params)

	u, err := uc.userInteractor.Create(appUser)
	if !errors.Is(err, nil) {
		return err
	}

	response := userToHTTPResponse(u)

	return c.JSON(http.StatusCreated, response)
}

// Mapping the business-oriented application model to the JSON http-bound structure
func userToHTTPResponse(user app.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

// Mapping the JSON http-bound structure to business-oriented application model
func httpRequestToUser(user User) app.User {
	return app.User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
