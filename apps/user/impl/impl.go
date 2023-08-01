package impl

import (
	"hezihua/apps"
	"hezihua/apps/user"
	"hezihua/conf"
	"sync"
)

type impl struct {
	Auth *conf.Auth
	Sessions map[string]string
	lock sync.Mutex
}

func (i *impl) Name() string {
	return user.AppName
}

func (i *impl) Init() error {
	i.Auth = conf.C().Auth
	i.Sessions = map[string]string{}
	return nil
}


func init(){
	apps.Registry(&impl{})
}