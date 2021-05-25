package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"golang-clean-architecture/config"
	"golang-clean-architecture/database"
	"golang-clean-architecture/registry"
	"golang-clean-architecture/router"
)

func main() {
	config.ReadConfig()

	db := database.NewDB()
	// db.LogMode(true)
	// defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
