package impl_test

import (
	"context"
	"hezihua/apps/user"
	"hezihua/apps/user/impl"
	"hezihua/test/tools"
	"os"
	"testing"
)

var (
	svc user.Service
	ctx = context.Background()
)


func TestLogin(t *testing.T) {
	req := user.NewLoginRequest()
	req.Username = os.Getenv("AUTH_USERNAME")
	req.Password = os.Getenv("AUTH_PASSWORD")
	ins, err := svc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestLogout(t *testing.T) {
	req := user.NewLogoutRequest()
	req.Username = os.Getenv("AUTH_USERNAME")
	err := svc.Logout(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init (){
	tools.DevelopmentSet()
	svc = impl.NewImpl()
}