package tests

import (
	"bytes"
	"github.com/alexperezortuno/toolbox/cmd/docker"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func TestDockerCmd(t *testing.T) {
	output, err := executeCommand(docker.DockerCmd)
	assert.NoError(t, err)
	assert.Contains(t, output, `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`)
}

func TestDockerCmdWithInvalidSubcommand(t *testing.T) {
	_, err := executeCommand(docker.DockerCmd, "invalid")
	assert.Error(t, err)
}

func TestDockerCmdWithImageSubcommand(t *testing.T) {
	_, err := executeCommand(docker.DockerCmd, "image")
	assert.NoError(t, err)
}

func TestDockerCmdWithVolumeSubcommand(t *testing.T) {
	_, err := executeCommand(docker.DockerCmd, "volume")
	assert.NoError(t, err)
}
