package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	saltCost := 12
	salt, err := bcrypt.GenerateFromPassword([]byte(password), saltCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(salt)
	return hashedPassword, nil
}

func ValidatePassword(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
