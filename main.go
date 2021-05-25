package main

import (
	"fmt"
	"log"

	"golang-clean-architecture/adapters/http"
	"golang-clean-architecture/adapters/repositories"
	"golang-clean-architecture/app"
	"golang-clean-architecture/config"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config.ReadConfig()

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	interactor := app.NewUserInteractor(repositories.NewUserRepository(db))
	controller := http.NewUserController(interactor)

	e := echo.New()
	e = http.NewRouter(e, controller)

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
