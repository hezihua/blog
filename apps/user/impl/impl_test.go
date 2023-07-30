package impl_test

import (
	"context"
	"hezihua/apps"
	"hezihua/apps/user"
	"hezihua/test/tools"
	"os"
	"testing"
)

var (
	svc user.Service
	ctx = context.Background()
)


func TestLogin(t *testing.T) {
	username := os.Getenv("AUTH_USERNAME")
	session := os.Getenv("AUTH_SESSION")
	req := user.NewLoginRequest()
	req.Username = username
	req.Password = os.Getenv("AUTH_PASSWORD")
	ins, err := svc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	checkreq := user.NewCheckIsLoginRequest(username, session)

	err = svc.CheckIsLogin(ctx, checkreq)
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

func TestIsLogin(t *testing.T) {
	username := os.Getenv("AUTH_USERNAME")
	session := os.Getenv("AUTH_SESSION")
	req := user.NewCheckIsLoginRequest(username, session)
	err := svc.CheckIsLogin(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init (){
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(user.AppName).(user.Service)
}

