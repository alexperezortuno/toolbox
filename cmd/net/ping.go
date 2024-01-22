package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

var (
	urlPath string
)

func ping(domain string) (string, error) {
	fmt.Println("Pinging...")
	var c *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		c = exec.Command("ping", "-n", "3", domain)
	case "linux":
		c = exec.Command("ping", "-c", "3", domain)
	case "darwin":
		c = exec.Command("ping", "-c", "3", domain)
	default:
		return "", fmt.Errorf("unsupported platform")
	}

	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
		return "", err
	}

	return string(out), nil
}

// pingCmd represents the ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping a host",
	Long:  `This command will ping a host and return the result`,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("\n%s\n", resp)
		}
	},
}

func init() {
	PingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to ping")

	if err := PingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
