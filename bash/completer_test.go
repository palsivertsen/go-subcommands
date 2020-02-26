package bash_test

import (
	"context"
	"os"
	"testing"

	base "github.com/palsivertsen/go-subcommands"
	"github.com/palsivertsen/go-subcommands/bash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComplete(t *testing.T) {
	tests := map[string]struct {
		compLine, compPoint string
		cmd                 base.Command
		expectedCompl       []string
	}{
		"sub completion": {
			compLine:      "rootCmd cm ignore this",
			compPoint:     "10", // cursor at "cm"
			expectedCompl: []string{"cmdA", "cmdB", "cmdC"},
			cmd: &command{
				subs: []base.Command{
					&command{name: "cmdA"},
					&command{name: "cmdC"},
					&command{name: "cmdB"},
					&command{name: "hello"},
					&command{name: "world"},
					&command{name: "command"},
				},
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			require.NoError(t, os.Setenv("COMP_LINE", tt.compLine))
			require.NoError(t, os.Setenv("COMP_POINT", tt.compPoint))

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			completions, err := bash.Complete(ctx, tt.cmd)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCompl, completions)
		})
	}
}

type command struct {
	base.UnimplementedCommand
	name string
	subs []base.Command
}

func (c *command) Name() string {
	return c.name
}

func (c *command) SubCommands() []base.Command {
	return c.subs
}
