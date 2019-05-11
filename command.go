package subcommands

import "flag"

type Command interface {
	Name() string
	SubCommands() []Command
	Flags() *flag.FlagSet
}

type UnimplementedCommand struct {
	Name_        string
	SubCommands_ []Command
	Flags_       flag.FlagSet
}

func (c *UnimplementedCommand) Name() string {
	return c.Name_
}

func (c *UnimplementedCommand) SubCommands() []Command {
	return c.SubCommands_
}

func (c *UnimplementedCommand) Flags() *flag.FlagSet {
	return &c.Flags_
}
