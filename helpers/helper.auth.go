package helpers

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/controllers/queries"

	"golang.org/x/crypto/bcrypt"
)

func IsRegistered(email string) (bool, error) {
	user, err := queries.GetUser(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	return true, nil
}

func HashPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(HashedPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}
