package file

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/spf13/cobra"
)

var (
	input  string
	output string
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress and decompress",
	Long:  `Compress and decompress files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

var tarCmd = &cobra.Command{
	Use:   "tar [flags]",
	Short: "Compress and decompress",
	Long:  `Compress and decompress files and directories with tar.`,
	Run: func(cmd *cobra.Command, args []string) {
		if input != "" && output != "" {
			c := "tar -czvf " + output + " " + input
			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}
			fmt.Printf(common.LineStr, out)
		} else {
			c := "tar -xvf " + input
			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}
			fmt.Printf(common.LineStr, out)
		}
	},
}

func init() {
	tarCmd.Flags().StringVarP(&input, "input", "i", "", "Input file or directory")
	tarCmd.Flags().StringVarP(&output, "output", "o", "", "Output file or directory")
	compressCmd.AddCommand(tarCmd)
}
