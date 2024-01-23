package common

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const LineStr string = "\n%s\n"
const CmdFailed string = "command failed with %s\n"
const PromptFailed string = "Prompt failed %v\n"

func LaunchCommand(command []string) (string, error) {
	c := exec.Command(command[0], command[1:]...)

	out, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stderr, label+" ")
		if err != nil {
			return ""
		}
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
