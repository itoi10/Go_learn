package gateways

import (
	"context"
	"errors"

	"example.com/entities"
	"example.com/usecases/ports"
)

// UserRepository実装

type UserGateway struct {
	// Todo
}

func NewUserRepository() ports.UserRepository {
	return &UserGateway{
		// Todo
	}
}

func (gateway *UserGateway) AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error) {
	// Todo
	return nil, errors.New("not implemented")
}

func (gateway *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	// Todo
	return nil, errors.New("not implemented")
}
