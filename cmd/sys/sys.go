package sys

import (
	"github.com/spf13/cobra"
)

var SysCmd = &cobra.Command{
	Use:   "sys",
	Short: "system management tools",
	Long:  `tools for system management`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {
	SysCmd.AddCommand(packageCmd)
}
