package flag_test

import (
	"testing"

	"github.com/revett/projects/internal/screenshot/flag"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestGlobalLength(t *testing.T) {
	require.Len(t, flag.Global(), 4)
}

func TestGlobalStringFlagsSingleAlias(t *testing.T) {
	for _, e := range flag.Global() {
		f, ok := e.(*cli.StringFlag)
		if !ok {
			continue
		}

		require.Len(t, f.Aliases, 1)
	}
}

func TestGlobalIntFlagsSingleAlias(t *testing.T) {
	for _, e := range flag.Global() {
		f, ok := e.(*cli.IntFlag)
		if !ok {
			continue
		}

		require.Len(t, f.Aliases, 1)
	}
}
