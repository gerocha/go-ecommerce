package factory

import (
    "tabajara/ecommerce/internal/domain/entity"
)

func NewClientFactory(id int, email string, name string, password string) (*client.ClientEntity, error) {
    return client.NewClientEntity(id, email, name, password)
}
