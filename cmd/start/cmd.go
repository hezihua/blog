package start

import (
	"hezihua/apps"
	"hezihua/conf"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	configType string
	filepath string
)

var RootCmd = &cobra.Command{
	Use: "start",
	Short: "vblog is a blog system",
	Long: "vblog is a blog system",
	Run: func(cmd *cobra.Command, args []string) {
		switch configType {
		case "file":
			err := conf.LoadConfigFromToml(filepath)
			cobra.CheckErr(err)
		case "env":
		default:
			cmd.PrintErr("unsupport config type")
		}

		// usrSvc := impl.NewImpl()
		err := apps.InitApps()
		cobra.CheckErr(err)

		server := gin.Default()
		v1 := server.Group("/vblog/api/v1")
		err = apps.InitHttpApps(v1)
		cobra.CheckErr(err)
	
		if err := server.Run(":8050"); err != nil {
			panic(err)
		}


		
	},
}


func init() {
	RootCmd.Flags().StringVarP(&configType, "config-type", "t", "file", "环境配置中加载")
	RootCmd.Flags().StringVarP(&filepath, "filepath", "f", "etc/config.toml", "配置文件路径")
}