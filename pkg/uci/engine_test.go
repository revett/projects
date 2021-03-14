package uci_test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

const mockEnginePath = "/path/to/engine"

func TestClose(t *testing.T) {
	uci.XCommand = mockCommander{}.Command
	e, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}

func TestNewEngine(t *testing.T) {
	uci.XCommand = mockCommander{}.Command
	_, err := uci.NewEngine(mockEnginePath)
	assert.NoError(t, err)
}

type mockCommander struct {
	out []string
}

func (m mockCommander) Command(s string, a ...string) *exec.Cmd {
	// nolint:gosec
	cmd := exec.Command(os.Args[0])
	out := fmt.Sprintf("TEST_CMD_OUTPUT=%s", strings.Join(m.out, ","))
	cmd.Env = append(os.Environ(), "TEST_MAIN=1", out)

	return cmd
}

func TestMain(m *testing.M) {
	if os.Getenv("TEST_MAIN") != "1" {
		os.Exit(m.Run())
	}

	l := strings.Split(os.Getenv("TEST_CMD_OUTPUT"), ",")
	for _, s := range l {
		fmt.Fprintln(os.Stdout, s)
		time.Sleep(1 * time.Millisecond)
	}

	os.Exit(0)
}
