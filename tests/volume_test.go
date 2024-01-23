package tests

import (
	"github.com/alexperezortuno/toolbox/cmd/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func VolumeCmdExecution(t *testing.T) {
	output, err := executeCommand(docker.VolumeCmd)
	assert.NoError(t, err)
	assert.Contains(t, output, "")
}

func RemoveVolumeCmdExecution(t *testing.T) {
	_, err := executeCommand(docker.VolumeCmd, "remove")
	assert.NoError(t, err)
}

func RemoveVolumeCmdWithAllFlagExecution(t *testing.T) {
	_, err := executeCommand(docker.VolumeCmd, "remove", "--all")
	assert.NoError(t, err)
}

func ListVolumeCmdExecution(t *testing.T) {
	_, err := executeCommand(docker.VolumeCmd, "list")
	assert.NoError(t, err)
}

func CreateVolumeCmdExecution(t *testing.T) {
	_, err := executeCommand(docker.VolumeCmd, "create")
	assert.NoError(t, err)
}
