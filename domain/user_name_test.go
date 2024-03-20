package domain_test

import (
	"domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserName(t *testing.T) {

	t.Run("success NewUserName()", func(t *testing.T) {
		name := "hogefuga"
		user_name, _ := domain.NewUserName(name)
		assert.Equal(t, name, user_name.ExportName())
	})

	t.Run("name length error NewUserName()", func(t *testing.T) {
		name := "hoge"
		_, err := domain.NewUserName(name)
		assert.Error(t, err)
	})
}
