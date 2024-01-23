package tests

import (
	"github.com/alexperezortuno/toolbox/cmd/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageCmd(t *testing.T) {
	output, err := executeCommand(docker.ImageCmd)
	assert.NoError(t, err)
	assert.Contains(t, output, "")
}

func TestListCmd(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "list")
	assert.NoError(t, err)
}

func TestPullCmd(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "pull")
	assert.NoError(t, err)
}

func TestBuildCmd(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "build")
	assert.NoError(t, err)
}

func TestHistoryCmd(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "history")
	assert.NoError(t, err)
}

func TestRemoveCmd(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "remove")
	assert.NoError(t, err)
}

func TestRemoveCmdWithAllFlag(t *testing.T) {
	_, err := executeCommand(docker.ImageCmd, "remove", "--all")
	assert.NoError(t, err)
}
