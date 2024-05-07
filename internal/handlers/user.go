package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"register/internal/database"
	"register/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GetAllUsers Get all user from Databse
//@Description Get all user inserted in Database
//@Tags Users
//@Produce json
//@Success 200 {array} models.GetUser
//Failure 500 {object} string "Failed to connect to database"
//Failure 500 {object} string "Failed to fetch users from Database"
//Failure 500 {object} string "Failed to scan rows from Database"
//@Router /getallusers [get]
func GetAllUsers(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect in Database")
		return c.JSON(http.StatusOK, "Failed to connect in Database")
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nome, email FROM users")
	if err != nil{
		log.Println("Failed to fetch users from Database")
		return c.JSON(http.StatusOK, "Failed to fetch users from Database")
	}
	defer rows.Close()

	var users []models.GetUser

	for rows.Next(){
		var user models.GetUser
		err = rows.Scan(&user.ID ,&user.Nome, &user.Email)
		if err != nil{
			log.Println("Failed to scan rows")
			return c.JSON(http.StatusOK ,"Failed to scan rows")
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

//GetUserByID Get user by id from Databse
//@Description Get a user by your ID in Database
//@Tags Users
//@Produce json
//@Param id path int true "User ID"
//@Security Bearer
//@Success 200 {array} models.GetUser
//Failure 500 {object} string "Failed to connect to database"
//Failure 500 {object} string "Failed to fetch users from Database"
//Failure 500 {object} string "Failed to scan rows from Database"
//@Router /getuserbyid/{id} [get]
func GetUserByID(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect in Database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect in Database")
	}
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	row := db.QueryRow("SELECT id, nome, email FROM users WHERE id = ?", userID)

	var user models.GetUser

	err = row.Scan(&user.ID, &user.Nome, &user.Email)
	if err != nil{
		if err == sql.ErrNoRows {
			log.Println("User not found")
			return c.JSON(http.StatusNotFound, "User not found")
		}
		log.Println("Failed to fetch user from Database")
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user from Database")
	}
	return c.JSON(http.StatusOK, user)
}

//UpdateUser Update user by id from Databse
//@Description Update a user by your ID in Database
//@Tags Users
//@Accept json
//@Produce json
//@Param id path int true "User ID to be updated"
//@Security Bearer
//@Success 200 {array} models.GetUser
//Failure 500 {object} string "Failed to connect to database"
//Failure 500 {object} string "Failed to decode Request Body"
//Failure 500 {object} string "Failed to update user"
//@Router /updateuser/{id} [put]
func UpdateUser(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect in Database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect in Database")
	}
	defer db.Close()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.GetUser
	err = c.Bind(&user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Failed to decode request body")
	}

	_, err = db.Exec("UPDATE users SET nome = ?, email= ? WHERE id = ?", user.Nome, user.Email, userID)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	updatedUser := models.GetUser{}
    err = db.QueryRow("SELECT id, nome, email FROM users WHERE id = ?", userID).Scan(&updatedUser.ID, &updatedUser.Nome, &updatedUser.Email)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to fetch updated user data")
    }
	return c.JSON(http.StatusOK, updatedUser)
}

//DeleteUser Delete user by ID from Databse
//@Description Delete a user by your ID in Database
//@Tags Users
//@Produce json
//@Param id path int true "User ID to be deleted"
//@Security Bearer
//@Success 200 {object} string "User deleted successfully"
//Failure 500 {object} string "Failed to connect to database"
//Failure 500 {object} string "Invalid user ID"
//Failure 500 {object} string "Failed to delete user"
//@Router /deleteuser/{id} [delete]
func DeleteUser(c echo.Context) error{
	db, err := database.Connect()
	if err != nil{
		log.Println("Failed to connect in Database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect in Database")
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Failed to delete user from databse")
	}
	return c.JSON(http.StatusOK, "User deleted succefully")
}