package domain

import (
	"errors"
	"unicode/utf8"
)

type UserName struct {
	name string
}

func NewUserName(name string) (*UserName, error) {
	if len := utf8.RuneCountInString(name); 6 > len || len > 20 {
		return nil, errors.New("name length error")
	}

	user_name := &UserName{}
	user_name.name = name

	return user_name, nil
}
