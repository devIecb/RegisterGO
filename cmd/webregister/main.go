package main

import (
	_ "register/api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//@title Web Register
//@description A Simple Register Web Service
//@host localhost:8001
//@basePath /
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8001"))
}