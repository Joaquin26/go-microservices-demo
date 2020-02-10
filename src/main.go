package main

import (
	"./routes"

	"github.com/labstack/echo"
)

//main function
func main() {
	// create a new echo instance
	e := echo.New()
	routes.SetTaskRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
