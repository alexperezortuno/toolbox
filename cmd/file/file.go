package file

import (
	"github.com/spf13/cobra"
)

var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "Download, compress, and decompress files and directories",
	Long:  `Tool for file operations such as downloading, compressing, and decompressing files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {
	FileCmd.AddCommand(downloadCmd)
	FileCmd.AddCommand(compressCmd)
}
