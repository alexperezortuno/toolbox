package cmd

import (
	"github.com/alexperezortuno/toolbox/cmd/docker"
	"github.com/alexperezortuno/toolbox/cmd/file"
	"github.com/alexperezortuno/toolbox/cmd/find"
	"github.com/alexperezortuno/toolbox/cmd/net"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	AddSubCommand()
}
