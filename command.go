package base

import "context"

// A Command is an item in a subcommand chain. Any command can be the root command in a chain as it knows nothing about it's parent.
// A command is allowed to configure it's children.
type Command interface {
	// Name of this command. Is used for walking the chain
	Name() string
	// Children of this command
	SubCommands() []Command
	Exec(context.Context, ...string) error
}

// RootCommand is a type one can embed into a struct to use as a base for other command implementations
type RootCommand struct {
	N    string
	Cmds []Command
}

// Name implements Command interface
func (c *RootCommand) Name() string {
	return c.N
}

// SubCommands implements Command interface
func (c *RootCommand) SubCommands() []Command {
	return c.Cmds
}

// Exec implements Command interface
func (c *RootCommand) Exec(ctx context.Context, args ...string) error {
	return nil
}
