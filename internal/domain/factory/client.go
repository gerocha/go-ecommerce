package factory

import (
    "tabajara/ecommerce/internal/domain/entity"
)

func NewClientFactory(id int, email string, name string, password string) (*client.Client, error) {
    return client.NewClientDomain(id, email, name, password)
}
