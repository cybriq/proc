package app

import (
	"testing"

	"github.com/cybriq/proc/pkg/cmds"
)

func TestNew(t *testing.T) {
	a, err := New(cmds.GetExampleCommands())
	_, _ = a, err
}
