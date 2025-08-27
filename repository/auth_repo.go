package repository

import (
	"context"
	"encoding/json"
	"test-asi/model"

	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	GetUser(username string) (*model.UserData, error)
	CreateUser(username string, userData *model.UserData) error
}

type userRepositoryImpl struct {
	client *redis.Client
	ctx    context.Context
}

func NewUserRepository(client *redis.Client) UserRepository {
	return &userRepositoryImpl{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *userRepositoryImpl) GetUser(username string) (*model.UserData, error) {
	key := "login_" + username
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var userData model.UserData
	if err := json.Unmarshal([]byte(val), &userData); err != nil {
		return nil, err
	}
	return &userData, nil
}

func (r *userRepositoryImpl) CreateUser(username string, userData *model.UserData) error {
	key := "login_" + username
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, key, userDataJSON, 0).Err()
}
