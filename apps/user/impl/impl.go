package impl

import (
	"hezihua/apps/user"
	"hezihua/conf"
	"sync"
)

func NewImpl() user.Service {
	return &impl{
		Auth: conf.C().Auth,
		Sessions: map[string]string{},
	}
}

type impl struct {
	Auth *conf.Auth
	Sessions map[string]string
	lock sync.Mutex
}