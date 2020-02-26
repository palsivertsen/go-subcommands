package base

import (
	"context"
	"flag"
	"strings"
)

type Command struct {
	// Name         string
	SubCommmands map[string]*Command
	Exec         func(ctx context.Context, args ...string) error
	Completer    Completer
}

type Completer interface {
	Complete(s string) []string
}

type PriorityCompleter struct {
	Completers []Completer
}

func (p *PriorityCompleter) Complete(s string) []string {
	for _, c := range p.Completers {
		if cs := c.Complete(s); len(cs) > 0 {
			return cs
		}
	}
	return nil
}

type FlagCompleter struct {
	FS *flag.FlagSet
}

func (f *FlagCompleter) Complete(s string) []string {
	args := strings.Split(s, " ")
	if len(args) == 0 {
		return nil
	}

	lArg := args[len(args)-1]

	var flags []string
	f.FS.VisitAll(func(f *flag.Flag) {
		tArg := strings.TrimPrefix(lArg, "-")
		if strings.HasPrefix(f.Name, tArg) {
			flags = append(flags, f.Name)
		}
	})
	return flags
}

type SubCommandCompleter struct {
	Subs map[string]*Command
}

func (scc *SubCommandCompleter) Complete(s string) []string {
	// Trim command
	args := strings.SplitN(s, " ", 3)

	switch len(args) {
	case 0, 1:
		return nil
	case 2:
		var r []string

		for k, _ := range scc.Subs {
			if strings.HasPrefix(k, args[1]) {
				r = append(r, k)
			}
		}
		return r
	}

	if c, ok := scc.Subs[args[1]]; ok && c.Completer != nil {
		return c.Completer.Complete(args[2])
	}

	return nil
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

// Exec implements Command interface
func (c *UnimplementedCommand) Exec(ctx context.Context, args ...string) error {
	return nil
}
