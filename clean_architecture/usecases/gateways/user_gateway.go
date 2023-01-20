package gateways

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	"example.com/entities"
	"example.com/usecases/ports"
)

// UserRepository実装
// firestoreに保存

type FIrestoreClientFactory interface {
	NewClient(ctx context.Context) (*firestore.Client, error)
}

type UserGateway struct {
	clientFactory FIrestoreClientFactory
}

func NewUserRepository(clientFactory FIrestoreClientFactory) ports.UserRepository {
	return &UserGateway{
		clientFactory: clientFactory,
	}
}

// ユーザデータをfirestoreに保存し,全てのユーザを返す
func (gateway *UserGateway) AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed AddUser NewClient: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{
		"age":     user.Age,
		"address": user.Address,
	})
	if err != nil {
		return nil, fmt.Errorf("failed AddUser Set: %v", err)
	}

	return getUsers(ctx, client)

}

// 全てのユーザを返す
func (gateway *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers NewClient: %v", err)
	}
	defer client.Close()

	return getUsers(ctx, client)
}

// firestoreから全てのユーザを取得する
func getUsers(ctx context.Context, client *firestore.Client) ([]*entities.User, error) {
	allData := client.Collection("users").Documents(ctx)

	docs, err := allData.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers GetAll: %v", err)
	}

	users := make([]*entities.User, 0)
	for _, doc := range docs {
		u := new(entities.User)
		err = mapToStruct(doc.Data(), &u)
		if err != nil {
			return nil, fmt.Errorf("failed GetUsers mapToStruct: %v", err)
		}
		u.Name = doc.Ref.ID
		users = append(users, u)
	}

	return users, nil
}

// mapを構造体に割り当てる
func mapToStruct(m map[string]interface{}, val interface{}) error {
	// mapをjsonへ変換
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// jsonをstructへ割り当て
	err = json.Unmarshal(buf, val)
	if err != nil {
		return err
	}

	return nil
}
