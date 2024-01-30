package net

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var (
	url string
)

var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "curl options",
	Long:  `Curl options for http requests and responses`,
	Run: func(cmd *cobra.Command, args []string) {

		if url == "" {
			switch runtime.GOOS {
			case "linux":
			case "darwin":
				executeCmd(cmd)
			default:
				fmt.Printf("unsupported platform")
				os.Exit(1)
			}
		}
	},
}

func executeCmd(cmd *cobra.Command) {
	if url == "" {
		fmt.Println("URL is required")
		err := cmd.Help()
		if err != nil {
			return
		}
		os.Exit(1)
	}

	var s []string
	s = append(s, "curl")
	s = append(s, url)
	out, err := common.LaunchCommand(s)
	if err != nil {
		fmt.Printf(common.CmdFailed, err)
	}
	fmt.Printf(common.LineStr, out)
}

func init() {
	curlCmd.Flags().StringVarP(&url, "url", "u", "", "url to request")
}
