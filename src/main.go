package main

import (
	"./entities"
	"./repository"
	"./routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

//main function
func main() {
	db := repository.Connect()
	db.AutoMigrate(&entities.Task{})
	defer db.Close()
	e := echo.New()
	routes.SetTaskRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
