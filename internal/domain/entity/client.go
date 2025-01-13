package client

import "tabajara/ecommerce/internal/domain/value_object"

type ClientEntity struct {
    Id int
    Email string
    Name string
    Password *value_object.Password
}

func NewClientEntity(id int, email string, name string, password string) (*ClientEntity, error) {
    pw, error := value_object.NewPassword(password)

    if error != nil {
        return nil, error
    }

    return &ClientEntity{id, email, name, pw}, nil
}
