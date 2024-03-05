package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
    password := RandomPassword()

    hashPassword, err := HashPassword(password)
    require.NoError(t, err)
    require.NotEmpty(t, hashPassword)

    err = ComparePassword(hashPassword, password)
    require.NoError(t, err)

    wrongPassword := RandomPassword()

    err = ComparePassword(hashPassword, wrongPassword)
    require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
