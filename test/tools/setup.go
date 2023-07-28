package tools

import (
	"hezihua/apps"
	_ "hezihua/apps/all"
	"hezihua/conf"
)

func DevelopmentSet(){
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	if err := apps.InitApps(); err != nil {
		panic(err)
	}

}