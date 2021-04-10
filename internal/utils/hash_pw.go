package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(userPassword string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {

	}

	return hash, err

}

func DecodePassword(userPassword string, hashpassword []byte) bool {

	if err := bcrypt.CompareHashAndPassword(hashpassword, []byte(userPassword)); err != nil {

		return false
	}

	return true
}
