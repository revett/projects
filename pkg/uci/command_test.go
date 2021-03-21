package uci_test

import (
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestIsReadyCommand(t *testing.T) {
	mc := mockCommander{
		out: []string{
			"readyok",
		},
	}

	e, err := uci.NewEngine(mc.Command, mockEnginePath)
	assert.NoError(t, err)

	err = e.Run(
		uci.IsReadyCommand(),
	)
	assert.NoError(t, err)

	err = e.Close()
	assert.NoError(t, err)
}
