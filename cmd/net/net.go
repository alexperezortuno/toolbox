/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"github.com/spf13/cobra"
)

// netCmd represents the net command
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
	// Here you will define your flags and configuration settings.
	NetCmd.AddCommand(PingCmd)
}
