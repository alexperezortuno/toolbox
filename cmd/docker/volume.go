package docker

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type volume struct {
	Name string
}

var VolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Commands for volumes",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

var removeVolumeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a volume or all volumes",
	Run: func(cmd *cobra.Command, args []string) {
		opt, _ := cmd.Flags().GetBool("all")
		if opt {
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to remove all volumes?",
				IsConfirm: true,
			}

			_, err := prompt.Run()

			if err != nil {
				fmt.Printf("Remove volumes cancel\n")
				os.Exit(0)
			}

			fmt.Printf("Removing all images\n")
			c := "docker volume rm $(docker volume ls -q)"
			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
				os.Exit(1)
			}
			fmt.Printf(common.LineStr, out)
			os.Exit(0)
		} else {
			c := "docker volume ls | awk 'NR>1{print $2}'"

			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf("command failed with %s\n", err)
			}

			if out == "" {
				fmt.Printf("No volumes to remove\n")
				os.Exit(0)
			}

			lines := strings.Split(out, "\n")
			var volumes []volume
			for _, line := range lines {
				if line == "" {
					continue
				}

				volumes = append(volumes, volume{
					Name: line,
				})
			}

			searcher := func(input string, index int) bool {
				vol := volumes[index]
				name := strings.Replace(strings.ToLower(vol.Name), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			}

			prompt := promptui.Select{
				Label:    "Select a volume",
				Items:    volumes,
				Size:     10,
				Searcher: searcher,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}?",
					Active:   "\U000025BB {{ .Name | cyan }}",
					Inactive: "  {{ .Name | cyan }}",
					Selected: "\U000025BB {{ .Name | red | cyan }}",
					Details: `
--------- Selected image ----------
{{ "Name:" | faint }}	{{ .Name }}`,
				},
			}

			i, _, err := prompt.Run()

			if err != nil {
				fmt.Printf(common.PromptFailed, err)
				return
			}

			out, err = common.LaunchCommand([]string{"docker", "volume", "rm", volumes[i].Name})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
				os.Exit(1)
			}

			fmt.Printf("Volume removed: %s", out)
		}
	},
}

var listVolumeCmd = &cobra.Command{
	Use:   "list",
	Short: "list all volumes",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := common.LaunchCommand([]string{"docker", "volume", "ls"})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("\n%s\n", out)
	},
}

var createVolumeCmd = &cobra.Command{
	Use:   "create",
	Short: "create a volume",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
			name = common.StringPrompt("Set volume name: ")
		}

		out, err := common.LaunchCommand([]string{"docker", "volume", "create", name})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("\nVolume created whit name: %s\n", out)
	},
}

func init() {
	removeVolumeCmd.Flags().Bool("all", false, "Remove all volumes")
	createVolumeCmd.Flags().StringP("name", "n", "", "Set volume name")
	VolumeCmd.AddCommand(listVolumeCmd)
	VolumeCmd.AddCommand(removeVolumeCmd)
	VolumeCmd.AddCommand(createVolumeCmd)
}
