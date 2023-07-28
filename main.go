package main

import (
	"hezihua/apps/api"
	"hezihua/apps/user/impl"
	"hezihua/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.LoadConfigFromToml("etc/config.toml")

	usrSvc := impl.NewImpl()


	userAPI := api.NewHandler(usrSvc)

	server := gin.Default()
  v1 := server.Group("/vblog/api/v1")
	userAPI.Registry(v1)

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}