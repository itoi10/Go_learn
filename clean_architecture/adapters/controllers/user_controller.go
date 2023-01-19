package controllers

import (
	"context"

	"example.com/usecases/ports"

	"github.com/labstack/echo/v4"
)

type User interface {
	AddUser(ctx context.Context) func(c echo.Context) error
	GetUsers(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInputPort
type RepositoryFactory func() ports.UserRepository
