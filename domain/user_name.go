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

type UserPass struct {
	pass string
}

func NewUserPass(pass string) (*UserPass, error) {
	if len := utf8.RuneCountInString(pass); 9 > len || len > 20 {
		return nil, errors.New("pass length error")
	}

	user_pass := &UserPass{}
	user_pass.pass = pass

	return user_pass, nil
}
