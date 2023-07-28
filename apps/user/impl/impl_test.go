package impl_test

import (
	"context"
	"hezihua/apps/user"
	"hezihua/apps/user/impl"
	"hezihua/test/tools"
	"testing"
)

var (
	svc user.Service
	ctx = context.Background()
)


func TestLogin(t *testing.T) {
	req := user.NewLoginRequest()
	ins, err := svc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func init (){
	tools.DevelopmentSet()
	svc = impl.NewImpl()
}