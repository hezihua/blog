package impl

import (
	"context"
	"fmt"
	"hezihua/apps/user"
)

func (i *impl) Login(ctx context.Context, req *user.LoginRequest) (*user.Session, error) {
	fmt.Println(i.Auth)
	return nil, nil
}

func (i *impl) Logout(ctx context.Context, req *user.LogoutRequest) error {
	return nil
}