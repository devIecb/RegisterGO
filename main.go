package main

import (
	_ "register/docs"
	"register/internal/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//@title Web Register
//@description A Simple Register Web Service
//@host localhost:8001
//@basePath /
func main() {
	e := echo.New()

	e.GET("/ping", handlers.Ping)
	e.GET("/getallusers", handlers.GetAllUsers)
	e.GET("/getuserbyid/:id", handlers.GetUserByID)
	e.POST("/signup", handlers.Signup)
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8001"))
}