//go:build wireinject
// +build wireinject

package drivers

// 依存関係の定義
// Wireを使ってDIコードを自動生成する

import (
	"context"

	"example.com/adapters/controllers"
	"example.com/adapters/presenters"
	"example.com/database"
	"example.com/usecases/gateways"
	"example.com/usecases/interactors"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeUserDriver(ctx context.Context) (User, error) {
	wire.Build(NewFirestoreClientFactory, echo.New, NewOutputFactory, NewInputFactory, NewRepositoryFactory, controllers.NewUserController, NewUserDriver)
	return &UserDriver{}, nil
}

func NewFirestoreClientFactory() gateways.FirestoreClientFactory {
	return &database.MyFirestoreClientFactory{}
}

func NewOutputFactory() controllers.OutputFactory {
	return presenters.NewUserOutputPort
}

func NewInputFactory() controllers.InputFactory {
	return interactors.NewUserInputPort
}

func NewRepositoryFactory() controllers.RepositoryFactory {
	return gateways.NewUserRepository
}
