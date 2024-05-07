package handlers

import (
	"log"
	"net/http"
	"register/internal/database"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error{
	db, err := database.Connect()
	if err != nil {
		log.Println("Failed to connect to database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping database")
		return c.JSON(http.StatusInternalServerError, "Failed to ping database")
	}

	return c.JSON(http.StatusOK, "all is working")
}
