package file

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var (
	urlPath    string
	outputPath string
	outputName string
	limit      int16
	retry      int16
	ftpUser    string
	ftpPass    string
	inputPath  string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file(s) from a URL",
	Long:  `For more information about this command, please visit: https://www.gnu.org/software/wget/manual/wget.html`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		var s []string

		background, _ := cmd.Flags().GetBool("background")
		if background {
			s = append(s, "-b")
		}

		cb, _ := cmd.Flags().GetBool("continue")
		if cb {
			s = append(s, "-c")
		}

		broken, _ := cmd.Flags().GetBool("broken")
		if broken {
			s = append(s, "-o wget-log -r -l 5 --spider")
		}

		download, _ := cmd.Flags().GetBool("download")
		if download {
			addS := []string{"--mirror", "--convert-links", "--page-requisites", "--no-parent", "-P", "documents/websites/"}
			s = append(s, addS...)
		}

		if ftpUser != "" {
			s = append(s, fmt.Sprintf("--ftp-user=%s", strings.TrimSpace(ftpUser)))
		}

		if ftpPass != "" {
			s = append(s, fmt.Sprintf("--ftp-password=%s", strings.TrimSpace(ftpPass)))
		}

		if inputPath != "" {
			s = append(s, fmt.Sprintf("-i %s", strings.TrimSpace(inputPath)))
		}

		if urlPath != "" {
			s = append(s, strings.TrimSpace(urlPath))
		}

		if outputPath != "" {
			s = append(s, fmt.Sprintf("-P %s", strings.TrimSpace(outputPath)))
		}

		if outputName != "" {
			s = append(s, fmt.Sprintf("-O %s", strings.TrimSpace(outputName)))
		}

		if limit != 0 {
			s = append(s, fmt.Sprintf("--limit-rate=%d", limit))
		}

		if retry != 0 {
			s = append(s, fmt.Sprintf("--tries=%d", retry))
		}

		c = exec.Command("wget", s...)
		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

func init() {
	downloadCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to download")
	downloadCmd.Flags().StringVarP(&outputName, "name", "n", "", "Output name")
	downloadCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path")
	downloadCmd.Flags().StringVarP(&ftpUser, "user", "U", "", "FTP user")
	downloadCmd.Flags().StringVarP(&ftpPass, "pass", "P", "", "FTP password")
	downloadCmd.Flags().Int16VarP(&limit, "limit", "l", 0, "Limit rate")
	downloadCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Input file")
	downloadCmd.Flags().Bool("background", false, "Run in background")
	downloadCmd.Flags().Bool("continue", false, "Continue")
	downloadCmd.Flags().Bool("broken", false, "Get broken links")
	downloadCmd.Flags().Bool("download", false, "Download site")
}
