package sub

import "flag"

// A Command is an item in a subcommand chain. Any command can be the root command in a chain as it knows nothing about it's parent.
// A command is allowed to configure it's children.
type Command interface {
	// Name of this command. Is used for walking the chain
	Name() string
	// Children of this command
	SubCommands() []Command
	// Flags for this particular command
	Flags() *flag.FlagSet
}

// UnimplementedCommand is a type one can embed into a struct to make it implement the Command interface without having to implement all the functions
type UnimplementedCommand struct {
}

// Name implements Command interface
func (c *UnimplementedCommand) Name() string {
	return ""
}

// SubCommands implements Command interface
func (c *UnimplementedCommand) SubCommands() []Command {
	return nil
}

// Flags implements Command interface
func (c *UnimplementedCommand) Flags() *flag.FlagSet {
	return nil
}
