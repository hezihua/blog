package main

import (
	_ "hezihua/apps/all"
	"hezihua/cmd"

	"github.com/spf13/cobra"
)

// func main() {
// 	conf.LoadConfigFromToml("etc/config.toml")

// 	// usrSvc := impl.NewImpl()
// 	err := apps.InitApps()
// 	if err != nil {
// 		panic(err)
// 	}

//   // arg0 := flag.Args()[0]
// 	// flag.Parse()

// 	// userAPI := api.NewHandler(usrSvc)

//   // v1 := server.Group("/vblog/api/v1")
// 	// userAPI.Registry(v1)
// 	server := gin.Default()
// 	v1 := server.Group("/vblog/api/v1")
// 	err = apps.InitHttpApps(v1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := server.Run(":8050"); err != nil {
// 		panic(err)
// 	}
// }


func main() {
	err := cmd.RootCmd.Execute()
	cobra.CheckErr(err)

}