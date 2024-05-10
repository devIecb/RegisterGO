package models

import "time"

type User struct {
	ID             int    `json:"id"`
	Nome           string    `json:"nome"`
	Email		   string 	 `json:"email"`
	Senha          string    `json:"senha"`
	Dt_cadastro    time.Time `json:"dt_cadastro"`
}

type UserLogin struct{
	Email 	string 	`json:"email"`
	Senha 	string 	`json:"senha"`
}

type GetUser struct {
	ID             int    	`json:"id"`
	Nome           string   `json:"nome"`
	Email		   string 	`json:"email"`
}
