package handlers

import (
	"log"
	"net/http"
	"register/internal/database"
	"register/internal/models"

	"github.com/labstack/echo/v4"
)

//Signup Create a new user
//@Description Create a new user using the provided informations
//@Tags Users
//@Accept json
//@Produce json
//@Param user body models.User true "User informations for registration"
//@Success 200 {object} models.User "Successfully create user"
//@Failure 500 {object} string "Failed to connect to database"
//@Failure 500 {object} string "Failed to decode request body"
//@Failure 500 {object} string "Failed to insert User in Database"
//@Router /signup [post]
func Signup(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect to database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer db.Close()

	var user models.User
	err = c.Bind(&user)
	if err != nil {
		log.Println("Failed to decode request body")
		return c.JSON(http.StatusInternalServerError, "Failed to decode request body")
	}

	_, err = db.Exec("INSERT INTO users (nome, email, senha) VALUES (?, ?, ?)", user.Nome, user.Email, user.Senha)
	if err != nil{
		log.Println("Failed to insert User in database")
		return c.JSON(http.StatusInternalServerError, "Failed to insert User in Database")
	}
	return c.JSON(http.StatusOK, "User created succesfully" )
}

func Login(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect in Database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect in Database")
	}
	defer db.Close()

	return nil
}
