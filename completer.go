package subcommands

import "context"

// Completer completes arguments
type Completer interface {
	// Complete returns suggested completions for the given command string
	Complete(ctx context.Context, cmd string) ([]string, error)
}
