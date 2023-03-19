package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := createTable()
	if err != nil {
		log.Fatal(err)
	}

	adminPass := initAdminUser()
	if adminPass != "" {
		println("Admin Password:", adminPass)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		return authUser(username, password), nil
	}))

	e.GET("/", indexRoute)
	e.GET("/:key", getRoute)
	e.POST("/:key", setRoute)
	e.DELETE("/:key", delRoute)

	e.POST("/users/:key", addUserRoute)
	e.DELETE("/users/:key", delUserRoute)

	e.Logger.Fatal(e.Start(":1323"))
}
