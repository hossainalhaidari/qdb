package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexRoute(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getRoute(c echo.Context) error {
	key := c.Param("key")

	val, err := get(key)
	if err != nil {
		if err.Error() == "not-found" {
			return c.NoContent(http.StatusNotFound)
		}

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, val)
}

func setRoute(c echo.Context) error {
	key := c.Param("key")
	val, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(string(val)) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	err = set(key, string(val))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func delRoute(c echo.Context) error {
	key := c.Param("key")

	err := del(key)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func addUserRoute(c echo.Context) error {
	key := c.Param("key")
	val, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(string(val)) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	err = addUser(key, string(val))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func delUserRoute(c echo.Context) error {
	key := c.Param("key")

	err := delUser(key)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
