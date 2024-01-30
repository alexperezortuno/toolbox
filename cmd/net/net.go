package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net command toolbox",
	Long:  `Net command toolbox`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {
	NetCmd.AddCommand(PingCmd)
	NetCmd.AddCommand(curlCmd)
}
