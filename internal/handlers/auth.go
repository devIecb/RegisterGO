package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"register/internal/database"
	"register/internal/models"
	"register/tools"

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

	Pass, err := tools.GenerateHash(user.Senha)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Error to generate hash password")
	}

	_, err = db.Exec("INSERT INTO users (nome, email, senha) VALUES (?, ?, ?)", user.Nome, user.Email, Pass)
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

	var user models.UserLogin
	err = c.Bind(&user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Error to decode request body")
	}

	var users models.User 
	err = db.QueryRow("SELECT email, senha FROM users WHERE email = ?", user.Email).Scan(&users.Email, &users.Senha)
	if err != nil{
		if err == sql.ErrNoRows{
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Email ou senha incorretass"})
		}
		return c.JSON(http.StatusInternalServerError, "Error to select users from Database")
	}
	
	comp := tools.CompareHashPassword(users.Senha, user.Senha)
	if !comp{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Email ou senha incorreta"})
	}
	return c.JSON(http.StatusOK, "User logged successfully")
}
		