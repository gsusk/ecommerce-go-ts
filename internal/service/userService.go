package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
}

func (svc UserService) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (svc UserService) Signup(params dto.UserSignUp) (string, error) {
	log.Panicln(params)
	return "", nil
}

func (svc UserService) Login(params any) (string, error) {
	return "", nil
}

func (svc UserService) VerificationCode(u domain.User) (int, error) {
	return 1, nil
}

func (svc UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (svc UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (svc UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (svc UserService) UpdateProfile(id uint, data any) error {
	return nil
}

func (svc UserService) BecomeSeller(id uint, data any) (string, error) {
	return "", nil
}

func (svc UserService) FindCart(id uint) ([]any, error) {
	return nil, nil
}

func (svc UserService) CreateCart(data any, user domain.User) (any, error) {
	return "", nil
}

func (svc UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (svc UserService) GetOrders(u domain.User) (any, error) {
	return nil, nil
}

func (svc UserService) GetOrderById(id uint, uId uint) (any, error) {
	return nil, nil
}
