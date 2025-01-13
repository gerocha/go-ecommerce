package client

import "tabajara/ecommerce/internal/domain/value_object"

type Client struct {
    Id int
    Email string
    Name string
    Password *value_object.Password
}

func NewClientDomain(id int, email string, name string, password string) (*Client, error) {
    pw, error := value_object.NewPassword(password)

    if error != nil {
        return nil, error
    }

    return &Client{id, email, name, pw}, nil
}
