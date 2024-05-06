package models

import "time"

type User struct {
	ID             string    `json:"id"`
	Nome           string    `json:"nome"`
	Email		   string 	 `json:"email"`
	Senha          string    `json:"senha"`
	Dt_cadastro    time.Time `json:"dt_cadastro"`
	Dt_exclusao    time.Time `json:"dt_exclusao"`
}

type UserLogin struct{
	Email 	string 	`json:"email"`
	Senha 	string 	`json:"senha"`
	Remember bool 	`json:"remember"`
}