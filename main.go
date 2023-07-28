package main

import (
	"hezihua/apps"
	"hezihua/conf"

	_ "hezihua/apps/all"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.LoadConfigFromToml("etc/config.toml")

	// usrSvc := impl.NewImpl()
	err := apps.InitApps()
	if err != nil {
		panic(err)
	}



	// userAPI := api.NewHandler(usrSvc)

	

	
  // v1 := server.Group("/vblog/api/v1")
	// userAPI.Registry(v1)
	server := gin.Default()
	v1 := server.Group("/vblog/api/v1")
	err = apps.InitHttpApps(v1)
	if err != nil {
		panic(err)
	}

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}