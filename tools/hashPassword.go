package tools

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(hash, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func GenerateHash(pass string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}