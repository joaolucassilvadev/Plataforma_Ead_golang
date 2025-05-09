package entities_test

import (
	"testing"

	"github.com/joaolucassilvadev/Plataforma_Ead_golang/service-core/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	t.Run("Should create a new user successfully", func(t *testing.T) {
		_, err := entities.NewUser("", "joao@gmail.com", "123456789", "12345678")

		assert.EqualError(t, err, "user name invalid")
	})
	t.Run("Should create a new user successfully", func(t *testing.T) {
		_, err := entities.NewUser("joao", "", "123456789", "12345678")

		assert.EqualError(t, err, "user email invalid")
	})

	t.Run("Should create a new user successfully", func(t *testing.T) {
		_, err := entities.NewUser("joao", "joao@gmail.com", "", "12345678")

		assert.EqualError(t, err, "user phone invalid")
	})

	t.Run("Should create a new user successfully", func(t *testing.T) {
		_, err := entities.NewUser("joao", "joao@gmail.com", "123456789", "")

		assert.EqualError(t, err, "Password must be at least 8 characters long")
	})

	t.Run("Should create a new user successfully", func(t *testing.T) {
		_, err := entities.NewUser("joao", "joao@gmail.com", "123456789", "123")

		assert.EqualError(t, err, "Password must be at least 8 characters long")
	})

	t.Run("Should create a new user successfully", func(t *testing.T) {
		user, err := entities.NewUser("joao", "joao@gmail.com", "123456789", "12345678")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotNil(t, user.Id)
		assert.Equal(t, user.Name, "joao")
		assert.Equal(t, user.Email, "joao@gmail.com")
		assert.Equal(t, user.Phone, "123456789")
		assert.Equal(t, user.Password, "12345678")
	})
}
