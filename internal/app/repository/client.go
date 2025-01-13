package repository

import (
    "tabajara/ecommerce/internal/domain/entity"
)

type IClientRepository interface {
    Add(client *client.ClientEntity)
}

type ClientRepository struct {
}

func (r ClientRepository) Add(client *client.ClientEntity) {
}
