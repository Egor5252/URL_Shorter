package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) ([]byte, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования")
	}

	return passHash, nil
}

func Compare(userPassHash string, incomingUserPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassHash), []byte(incomingUserPass))
	return err
}
