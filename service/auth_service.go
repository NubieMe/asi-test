package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"test-asi/model"
	"test-asi/repository"
)

type AuthService interface {
	Register(payload *model.RegisterPayload) error
	Login(payload *model.LoginPayload) (*model.UserData, error)
	HashPassword(password string) string
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) HashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *authService) Register(payload *model.RegisterPayload) error {
	existingUser, err := s.userRepo.GetUser(payload.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already registered")
	}

	userData := &model.UserData{
		Realname: payload.Realname,
		Email:    payload.Email,
		Password: s.HashPassword(payload.Password),
	}

	return s.userRepo.CreateUser(payload.Username, userData)
}

func (s *authService) Login(payload *model.LoginPayload) (*model.UserData, error) {
	userData, err := s.userRepo.GetUser(payload.Username)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errors.New("username/password is incorrect")
	}

	if s.HashPassword(payload.Password) != userData.Password {
		return nil, errors.New("username/password is incorrect")
	}

	return userData, nil
}
