package base

import "context"

// A Command is an item in a subcommand chain. Any command can be the root command in a chain as it knows nothing about it's parent.
// A command is allowed to configure it's children.
type Command interface {
	// Name of this command. Is used for walking the chain
	Name() string
	// Children of this command
	SubCommands() []Command
	// Known flags for the current command
	Flags() []string
	Exec(context.Context, ...string) error
}
