package user

import (
	"context"
	"net/http"
)

var (
	AppName = "user"
)

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func NewLoginRequestFromBasicAuth(r *http.Request) *LoginRequest {
	user, pass, _ := r.BasicAuth()
	return &LoginRequest{
		Username: user,
		Password: pass,
	}
}

type LoginRequest struct {
	Username string
	Password string
}

func NewLogoutRequest() *LogoutRequest {
	return &LogoutRequest{}
}

type LogoutRequest struct {
	Username string `json:"username"`
}

type Service interface {
	Login(context.Context, *LoginRequest) (*Session, error)

	Logout (context.Context, *LogoutRequest) error 

	CheckIsLogin(context.Context, *CheckIsLoginRequest)  error
}

func NewCheckIsLoginRequest(Username string, session string) *CheckIsLoginRequest {
	return &CheckIsLoginRequest{
		Username: Username,
		SessionId: session,
	}
}

type CheckIsLoginRequest struct {
	Username string
	SessionId string
}


