package initial

import "github.com/spf13/cobra"

var (
	version bool
)

var RootCmd = &cobra.Command{
	Use: "init",
	Short: "vblog is a blog system",
	Long: "vblog is a blog system",
	Run: func(cmd *cobra.Command, args []string) {
		// if version {
		// 	cmd.Printf("vblog-api version: %s\n", "v0.0.1")
		// 	return
		// }
		cmd.Help()
	},
}


func init() {
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "show version")
}