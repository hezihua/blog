package cmd

import (
	"hezihua/cmd/initial"
	"hezihua/cmd/start"
	"hezihua/version"

	"github.com/spf13/cobra"
)

var (
	showVersion bool
)

var RootCmd = &cobra.Command{
	Use: "vblog-api",
	Short: "vblog is a blog system",
	Long: "vblog is a blog system",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			// cmd.Printf("vblog-api version: %s\n", "v0.0.1")
			cmd.Printf("vblog-api version: %s\n", version.FullVersion())
			return
		}
		cmd.Help()
	},
}


func init() {

	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")

	RootCmd.AddCommand(start.RootCmd, initial.RootCmd)
}