package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"theaveasso.bab/internal/utility"
)

func createRandomUser(t *testing.T) User {
    userName := utility.RandomUsername()
    arg := CreateUserParams{
        Username       : userName,
        HashedPassword: utility.RandomPassword(), 
        FullName:  utility.RandomFullName(),
        Email: utility.RandomEmail(userName),
    }

    user, err := testQueries.CreateUser(context.Background(), arg)
    
    require.NoError(t, err)
    require.NotEmpty(t, user)

    require.Equal(t, arg.Email, user.Email)
    require.Equal(t, arg.Username, user.Username)
    require.Equal(t, arg.FullName, user.FullName)
    require.Equal(t, arg.HashedPassword, user.HashedPassword)
    require.NotZero(t, user.CreatedAt)
    require.True(t, user.PasswordChangedAt.IsZero())

    return user
}

func TestCreateUser(t *testing.T) {
    createRandomUser(t)
}

func TestGetUser(t *testing.T) {
    userCreated := createRandomUser(t)
    userGetted, err := testQueries.GetUser(context.Background(), userCreated.Username)

    require.NoError(t, err)
    require.NotEmpty(t, userGetted)


    require.Equal(t, userCreated.Email, userGetted.Email)
    require.Equal(t, userCreated.Username, userGetted.Username)
    require.Equal(t, userCreated.FullName, userGetted.FullName)
    require.Equal(t, userCreated.HashedPassword, userGetted.HashedPassword)
    require.WithinDuration(t, userCreated.CreatedAt, userGetted.CreatedAt, time.Second)
    require.WithinDuration(t, userCreated.PasswordChangedAt, userGetted.PasswordChangedAt, time.Second)
}
