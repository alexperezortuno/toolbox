package sys

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Package management tools",
	Long:  `Options for package management`,
	Run: func(cmd *cobra.Command, args []string) {
		var s []string
		var out string
		var err error
		s = append(s, "cat")
		s = append(s, "/etc/*-release")
		fmt.Printf(common.LineStr, strings.Join(s, " "))

		switch runtime.GOOS {
		case "linux":
			linuxOs(cmd, &s, &out, &err)
		default:
			err = fmt.Errorf("unsupported platform")
		}
	},
}

func linuxOs(cmd *cobra.Command, s *[]string, out *string, err *error) {
	common.OsDetail()
	*out, *err = common.LaunchCommand([]string{"bash", "-c", strings.Join(*s, " ")})

	if err != nil {
		fmt.Printf(common.CmdFailed, *err)
	}

	fmt.Printf(common.LineStr, *out)
	var id = common.ParseDistro(*out, "ID=")
	fmt.Printf(common.LineStr, id)

	opt, _ := cmd.Flags().GetBool("upgrade")
	if opt {
		pck, _ := cmd.Flags().GetString("package")

		if pck == "" {
			*out, *err = common.LaunchCommand([]string{"sudo", "apt-get", "install", pck, "--only-upgrade"})
			if *err != nil {
				fmt.Printf(common.CmdFailed, *err)
				os.Exit(1)
			}
			fmt.Printf(common.LineStr, *out)
			os.Exit(0)
		} else {
			fmt.Printf("Package name is required\n")
			os.Exit(1)
		}
	}

	opt, _ = cmd.Flags().GetBool("list")
	if opt {
		//sudo apt list --installed | less
		*out, *err = common.LaunchCommand([]string{"sudo", "apt", "list", "--installed"})
		if *err != nil {
			fmt.Printf(common.CmdFailed, *err)
			os.Exit(1)
		}
		fmt.Printf(common.LineStr, *out)
		os.Exit(0)
	}
}

func init() {
	packageCmd.Flags().StringP("package", "p", "", "Package name")
	packageCmd.Flags().BoolP("upgrade", "u", false, "Upgrade package")
	packageCmd.Flags().BoolP("list", "l", false, "Package list")
	packageCmd.Flags().Bool("less", false, "Use less")
}
