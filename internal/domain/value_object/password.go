package value_object

import (
	"errors"
)

type Password struct {
    Value string
}

func (password Password) IsValidPassword() (bool, error) {
    if len(password.Value) < 4 {
        return false, errors.New("Tamanho da senha pequeno")
    }
    return true, nil;
}

func NewPassword(password string) (*Password, error) {
    pw := &Password{password}
    is_valid, error := pw.IsValidPassword()
    if !is_valid {
        return nil, error
    }
    return &Password{password}, nil
}
