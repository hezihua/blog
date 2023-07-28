package impl

import (
	"hezihua/apps/user"
	"hezihua/conf"
)

func NewImpl() user.Service {
	return &impl{
		Auth: conf.C().Auth,
	}
}

type impl struct {
	Auth *conf.Auth
}