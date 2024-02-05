package cmd

import (
	"github.com/alexperezortuno/toolbox/cmd/docker"
	"github.com/alexperezortuno/toolbox/cmd/file"
	"github.com/alexperezortuno/toolbox/cmd/find"
	"github.com/alexperezortuno/toolbox/cmd/net"
	"github.com/alexperezortuno/toolbox/cmd/sys"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func AddSubCommand() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(file.FileCmd)
	rootCmd.AddCommand(find.FinderCmd)
	rootCmd.AddCommand(sys.SysCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	AddSubCommand()
}
