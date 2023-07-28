package apps

import (
	"hezihua/logger"

	"github.com/gin-gonic/gin"
)

var (
	httpAppStore = map[string]HttpIocObject{}
)

type HttpIocObject interface {
	IocObject
	RegistryHandler(gin.IRouter)
}


func RegistryHttp(obj HttpIocObject) {
	if _, ok := httpAppStore[obj.Name()]; ok {
		panic("ioc object has been registry")
	}
	httpAppStore[obj.Name()] = obj
	logger.L().Debug().Msgf("http registry ioc object: %s", obj.Name())
}

func GetHttpApp(objName string) any {
	if v, ok := httpAppStore[objName]; ok {
		return v
	}
	panic("ioc object not found")
}