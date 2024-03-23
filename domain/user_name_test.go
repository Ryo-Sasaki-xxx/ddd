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

	t.Run("success NewUserPass()", func(t *testing.T) {
		pass := "hogehogehoge"
		user_pass, _ := domain.NewUserPass(pass)
		assert.Equal(t, pass, user_pass.ExportPass())
	})

	t.Run("pass length error NewUserPass()", func(t *testing.T) {
		pass := "hoge"
		_, err := domain.NewUserPass(pass)
		assert.Error(t, err)
	})
}
