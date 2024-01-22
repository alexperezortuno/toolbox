/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package docker

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var (
	list bool
)

// imageCmdCmd represents the imageCmd command
var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Commands for images",
	Long: `This is a compilated commands for images For example:

toolbox docker image -l`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all images",
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		c = exec.Command("docker", "image", "ls")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf("command failed with %s\n", err)
		}
		fmt.Printf("\n%s\n", out)
	},
}

func init() {
	ImageCmd.AddCommand(listCmd)
	//ImageCmd.Flags().BoolVarP(&list, "list", "l", false, "List all images")
}
