package handler

import (
	"echo_sample/model"
	"net/http"

	"github.com/labstack/echo"
)

// 新規Todo登録
func AddTodo(c echo.Context) error {
	// パラメータの値を構造体に格納
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	// フィールドチェック
	if todo.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid todo name fields",
		}
	}

	// JWT内のユーザーIDがDBに存在するか
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	// TodoをDBに登録
	todo.UID = uid
	model.CreateTodo(todo)

	// 登録したTodoを返却
	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	// JWT内のユーザーIDがDBに存在するか
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todos := model.FindTodos(&model.Todo{UID: uid})
	return c.JSON(http.StatusOK, todos)
}
