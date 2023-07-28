package tools

import "hezihua/conf"

func DevelopmentSet(){
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
}