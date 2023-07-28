package apps

import (
	"hezihua/logger"
)

var (
	internalAppStore = map[string]IocObject{}
)

type IocObject interface {
	Name() string
	Init() error
}


func Registry(obj IocObject) {
	if _, ok := internalAppStore[obj.Name()]; ok {
		panic("ioc object has been registry")
	}
	internalAppStore[obj.Name()] = obj
	logger.L().Debug().Msgf("registry ioc object: %s", obj.Name())
}

func GetInternalApp(objName string) any {
	if v, ok := internalAppStore[objName]; ok {
		return v
	}
	panic("ioc object not found")
}