package bash_test

import (
	"context"
	"os"
	"testing"

	"github.com/palsivertsen/go-subcommands"
	"github.com/palsivertsen/go-subcommands/bash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComplete(t *testing.T) {
	tests := map[string]struct {
		compLine, compPoint string
		cmd                 subcommands.Command
		expectedCompl       []string
	}{
		"sub completion": {
			compLine:      "rootCmd cm ignore this",
			compPoint:     "10", // cursor at "cm"
			expectedCompl: []string{"cmdA", "cmdB", "cmdC"},
			cmd: &subcommands.UnimplementedCommand{
				SubCommands_: []subcommands.Command{
					&subcommands.UnimplementedCommand{Name_: "cmdA"},
					&subcommands.UnimplementedCommand{Name_: "cmdC"},
					&subcommands.UnimplementedCommand{Name_: "cmdB"},
					&subcommands.UnimplementedCommand{Name_: "hello"},
					&subcommands.UnimplementedCommand{Name_: "world"},
					&subcommands.UnimplementedCommand{Name_: "command"},
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
