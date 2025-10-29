package service

import "go-ecommerce-app/internal/domain"

type UserService struct {
}

func (svc UserService) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (svc UserService) Signup(params any) (string, error) {
	return "", nil
}

func (svc UserService) Login(params any) (string, error) {
	return "", nil
}

func (svc UserService) VerificationCode(params any) (string, error) {
	return "", nil
}

func (svc UserService) VerifyCode(params any) (string, error) {
	return "", nil
}

func (svc UserService) CreateProfile(params any) (string, error) {
	return "", nil
}

func (svc UserService) GetProfile(params any) (string, error) {
	return "", nil
}

func (svc UserService) UpdateProfile(params any) (string, error) {
	return "", nil
}

func (svc UserService) BecomeSeller(params any, data any) (string, error) {
	return "", nil
}

func (svc UserService) FindCart(id uint) (any, error) {
	return "", nil
}

func (svc UserService) CreateCart(data any, user domain.User) (string, error) {
	return "", nil
}

func (svc UserService) CreateOrder(u domain.User) (string, error) {
	return "", nil
}

func (svc UserService) GetOrders(u domain.User) (any, error) {
	return nil, nil
}

func (svc UserService) GetOrderById(id uint) (any, error) {
	return nil, nil
}
