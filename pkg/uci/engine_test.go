package uci_test

import (
	"testing"

	"github.com/revett/projects/pkg/uci"
	"github.com/stretchr/testify/assert"
)

func TestEngineIsReady(t *testing.T) {
	e, err := uci.NewEngine("/usr/local/bin/stockfish")
	assert.NoError(t, err)

	ready, err := e.IsReady()
	assert.NoError(t, err)
	assert.True(t, ready)

	err = e.Stop()
	assert.NoError(t, err)
}
