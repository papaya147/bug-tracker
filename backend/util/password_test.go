package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPassword(t *testing.T) (string, string) {
	password := RandomString(20)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	return password, hashedPassword
}

func TestHashPassword(t *testing.T) {
	createRandomPassword(t)
}

func TestValidatePassword(t *testing.T) {
	password, hashedPassword := createRandomPassword(t)

	err := ValidatePassword(password, hashedPassword)
	require.NoError(t, err)
}
