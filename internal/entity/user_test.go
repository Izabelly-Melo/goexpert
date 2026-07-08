package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Lena", "lena@example.com", "12345678")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Lena", user.Name)
	assert.Equal(t, "lena@example.com", user.Email)
}

func TestCheckPassword(t *testing.T) {
	user, err := NewUser("Lena", "lena@example.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("123456"))
	assert.False(t, user.CheckPassword("wrongpassword"))
	assert.NotEqual(t, "12345678", user.Password)
}
