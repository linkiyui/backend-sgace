package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func MainHashingFunc() {
	password := "samplePassword"
	hashedPassword := HashingPasswordFunc(password)

	if hashedPassword == "" {
		panic("Error while hashing password")
	}
	fmt.Println("Hashed Password:", hashedPassword)
	fmt.Println("Password Match:", CheckPasswordHashFunc(password, hashedPassword))
}

func HashingPasswordFunc(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

		panic(err)
	}
	return string(hash)
}

func CheckPasswordHashFunc(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
