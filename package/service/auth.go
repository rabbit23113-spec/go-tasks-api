package service

import "main/package/types"

type Authentication interface {
	SignUp(form types.SignUpDto) error
	SignIn(form types.SignInDto) error
}

type AuthService struct{}

func (a *AuthService) SignUp(form types.SignUpDto) error {
	return nil
}

func (a *AuthService) SignIn(form types.SignInDto) error {
	return nil
}
