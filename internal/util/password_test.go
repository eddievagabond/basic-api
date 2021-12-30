package util_test

import (
	"testing"

	"github.com/eddievagabond/internal/util"
	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "testPassword"

	// Ensure the password is hashed correctly
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// Ensure validating correct password succeeds
	err = util.CheckPasswordHash(password, hashedPassword)
	require.NoError(t, err)

	// Ensure invalid password fails
	wrongPassword := "wrongPassword"
	err = util.CheckPasswordHash(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
