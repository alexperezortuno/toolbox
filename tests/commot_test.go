package tests_test

import (
	"errors"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCommand struct {
	mock.Mock
}

func (m *MockCommand) CombinedOutput() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func TestLaunchCommandSuccessfulExecution(t *testing.T) {
	mockCmd := new(MockCommand)
	mockCmd.On("CombinedOutput").Return([]byte("success"), nil)

	common.CommandCreator = func(name string, arg ...string) common.Cmd {
		return mockCmd
	}

	output, err := common.LaunchCommand([]string{"ls", "-la"})
	assert.NoError(t, err)
	assert.Equal(t, "success", output)
}

func TestLaunchCommandFailedExecution(t *testing.T) {
	mockCmd := new(MockCommand)
	mockCmd.On("CombinedOutput").Return([]byte{}, errors.New("command failed"))

	common.CommandCreator = func(name string, arg ...string) common.Cmd {
		return mockCmd
	}

	_, err := common.LaunchCommand([]string{"invalid", "command"})
	assert.Error(t, err)
}

func TestLaunchCommandEmptyCommand(t *testing.T) {
	mockCmd := new(MockCommand)
	mockCmd.On("CombinedOutput").Return(nil, errors.New("command failed"))

	common.CommandCreator = func(name string, arg ...string) common.Cmd {
		return mockCmd
	}

	_, err := common.LaunchCommand([]string{})
	assert.Error(t, err)
}
