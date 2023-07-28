package apps

import (
	"hezihua/logger"

	"github.com/gin-gonic/gin"
)

func InitApps() error {
	for _, v := range internalAppStore {
		if err := v.Init(); err != nil {
			panic(err)
		}
		logger.L().Debug().Msgf("init ioc object: %s", v.Name())
	}
	return nil
}

func InitHttpApps(r gin.IRouter) error {
	for _, v := range httpAppStore {
		if err := v.Init(); err != nil {
			panic(err)
		}
		v.RegistryHandler(r.Group(v.Name()))
		logger.L().Debug().Msgf("init ioc object: %s", v.Name())
	}
	return nil
}