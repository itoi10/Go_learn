package controllers

import (
	"context"

	"example.com/entities"
	"example.com/usecases/gateways"
	"example.com/usecases/ports"

	"github.com/labstack/echo/v4"
)

// InputPort, OutputPort, Repositoryを組み立ててInputPortを実行する

type User interface {
	AddUser(ctx context.Context) func(c echo.Context) error
	GetUsers(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.UserOutputPort
type InputFactory func(ports.UserOutputPort, ports.UserRepository) ports.UserInputPort
type RepositoryFactory func(gateways.FirestoreClientFactory) ports.UserRepository

type UserController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	clientFactory     gateways.FirestoreClientFactory
}

func NewUserController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory, clientFactory gateways.FirestoreClientFactory) User {
	return &UserController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
		clientFactory:     clientFactory,
	}
}

func (u *UserController) AddUser(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		user := new(entities.User)
		if err := c.Bind(user); user != nil {
			return err
		}

		return u.newInputPort(c).AddUser(ctx, user)
	}
}

func (u *UserController) GetUsers(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		return u.newInputPort(c).GetUsers(ctx)
	}
}

func (u *UserController) newInputPort(c echo.Context) ports.UserInputPort {
	outputPort := u.outputFactory(c)
	repository := u.repositoryFactory(u.clientFactory)
	return u.inputFactory(outputPort, repository)
}
