package ports

import (
	"context"

	"example.com/entities"
)

// ユーザデータ受け取り
type UserInputPort interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUsers(ctx context.Context) error
}

// ユーザデータを返す
type UserOutputPort interface {
	OutputUsers([]*entities.User) error
	OutputError(error) error
}

// データ保存
type UserRepository interface {
	AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error)
	GetUsers(ctx context.Context) ([]*entities.User, error)
}
