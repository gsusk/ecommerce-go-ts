package service

import (
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (svc UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := svc.Repo.FindUser(email)
	return &user, err
}

func (svc UserService) Signup(params dto.UserSignUp) (string, error) {
	user, err := svc.Repo.CreateUser(domain.User{
		Email:    params.Email,
		Password: params.Password,
		Phone:    params.Phone,
	})

	if err != nil {
		return "", fmt.Errorf("error creating user: %v", err)
	}

	userInfo := fmt.Sprintf("%v, %v, %v", user.Email, user.ID, user.UserType)

	return userInfo, nil
}

func (svc UserService) Login(email string, password string) (string, error) {
	user, err := svc.Repo.FindUser(email)
	if err != nil {
		return "", err
	}
	userInfo := fmt.Sprintf("%v, %v, %v", user.Email, user.ID, user.UserType)

	return userInfo, nil
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
