package entity

import "tabajara/ecommerce/internal/domain/value_object"

type Account struct {
    Id int
    Email string
    Password value_object.Password
}

func New(id int, email string, password string) *Account {
    var pw value_object.Password
    pw.Value = password

    return &Account{id, email, pw}
}
