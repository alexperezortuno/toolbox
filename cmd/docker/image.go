/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package docker

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var (
	url string
	tag string
	all bool
)

type image struct {
	Name string
	Id   string
	Tag  string
}

var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Commands for images",
	Long:  `This command is used for managing images`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all images",
	Long: `This command is used to listing images
For example:
toolbox docker image list`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		c = exec.Command("docker", "image", "ls")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image or a repository from a registry",
	Long: `This command is used to pull an image or a repository from a registry
For example:
toolbox docker image pull`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		c = exec.Command("docker", "image", "pull")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an image from a Dockerfile or url",
	Long: `This command is used for building from the current directory’s dockerfile or a url
For example:
toolbox docker image build`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		c = exec.Command("docker", "image", "build", url)

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Showing the History of an Image",
	Long: `This command will let know the history of the image inside the docker
For example:
toolbox docker image history`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd
		c = exec.Command("docker", "image", "history", "atmoz/sftp:latest")

		out, err := c.CombinedOutput()
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove one or more images",
	Long: `This command is used to remove one or more images
For example:
toolbox docker image remove -t latest`,
	Run: func(cmd *cobra.Command, args []string) {
		opt, _ := cmd.Flags().GetBool("all")
		if opt {
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to remove all images?",
				IsConfirm: true,
			}

			_, err := prompt.Run()

			if err != nil {
				fmt.Printf("Remove images cancel\n")
				os.Exit(0)
			}

			fmt.Printf("Removing all images\n")

			c := "docker stop $(docker ps -a -q)"
			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}
			fmt.Printf(common.LineStr, out)

			c = "docker rm $(docker ps -a -q)"
			out, err = common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}
			fmt.Printf(common.LineStr, out)

			c = "docker rmi $(docker images -q)"
			out, err = common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}
			fmt.Printf(common.LineStr, out)
		} else {
			c := "docker image list | awk 'NR>1{print $1\"|\"$3\"|\"$2}'"

			out, err := common.LaunchCommand([]string{"bash", "-c", c})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
			}

			lines := strings.Split(out, "\n")
			var images []image
			for _, line := range lines {
				if line == "" {
					continue
				}
				img := strings.Split(line, "|")
				images = append(images, image{
					Name: img[0],
					Id:   img[1],
					Tag:  img[2],
				})
			}

			searcher := func(input string, index int) bool {
				pepper := images[index]
				name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			}

			prompt := promptui.Select{
				Label:    "Select Day",
				Items:    images,
				Size:     10,
				Searcher: searcher,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}?",
					Active:   "\U000025BB {{ .Name | cyan }} ({{ .Id | red }})",
					Inactive: "  {{ .Name | cyan }} ({{ .Id | red }})",
					Selected: "\U000025BB {{ .Name | red | cyan }}",
					Details: `
--------- Selected image ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Id:" | faint }}	{{ .Id }}`,
				},
			}

			i, _, err := prompt.Run()

			if err != nil {
				fmt.Printf(common.PromptFailed, err)
				return
			}

			out, err = common.LaunchCommand([]string{"docker", "image", "rm", images[i].Id})
			if err != nil {
				fmt.Printf(common.CmdFailed, err)
				os.Exit(1)
			}

			fmt.Printf("\n%s\n", out)
		}
	},
}

func init() {
	buildCmd.Flags().StringVarP(&url, "url", "u", ".", "Url path")
	ImageCmd.AddCommand(listCmd)
	ImageCmd.AddCommand(pullCmd)
	ImageCmd.AddCommand(buildCmd)
	ImageCmd.AddCommand(historyCmd)
	ImageCmd.AddCommand(removeCmd)
	historyCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "Tag name")
	removeCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "Tag name")
	removeCmd.Flags().Bool("all", false, "Remove all images")
}
