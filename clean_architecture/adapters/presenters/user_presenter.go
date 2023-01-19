package presenters

import (
	"log"
	"net/http"

	"example.com/entities"
	"example.com/usecases/ports"
	"github.com/labstack/echo/v4"
)

// OutputPortの実装
// echoでユーザデータのJSON or エラーを返す

type UserPresenter struct {
	ctx echo.Context
}

func NewUserOutputPort(ctx echo.Context) ports.UserOutputPort {
	return &UserPresenter{
		ctx: ctx,
	}
}

func (u *UserPresenter) OutputUsers(users []*entities.User) error {
	return u.ctx.JSON(http.StatusOK, users)
}

func (u *UserPresenter) OutputError(err error) error {
	log.Fatal(err)
	return u.ctx.JSON(http.StatusInternalServerError, err)
}
