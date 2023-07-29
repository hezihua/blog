package impl

import (
	"context"
	"hezihua/apps/user"
)

func (i *impl) Login(ctx context.Context, req *user.LoginRequest) (*user.Session, error) {
	if err := i.Auth.Validate(req.Username, req.Password); err != nil {
		return nil, err
	}
	// 记录一个session
	sess := i.createSession(req.Username)
	return user.NewSession(req.Username, sess), nil
}

func (i *impl) Logout(ctx context.Context, req *user.LogoutRequest) error {
	i.deleteSession(req.Username)
	return nil
}