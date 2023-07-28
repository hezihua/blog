package user

import "context"

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LoginRequest struct {
	username string
	password string
}

func NewLogoutRequest() *LogoutRequest {
	return &LogoutRequest{}
}

type LogoutRequest struct {
 
}

type Service interface {
	Login(context.Context, *LoginRequest) (*Session, error)

	Logout (context.Context, *LogoutRequest) error 
}

