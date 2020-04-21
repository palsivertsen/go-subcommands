package subcommands

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type SubCommandExecutor map[string]Executor

func (subex SubCommandExecutor) Exec(ctx context.Context, args ...string) error {
	if len(args) <= 1 {
		return errors.New("no sub command")
	}

	ex, ok := subex[args[1]]
	if !ok {
		return fmt.Errorf("unknown sub command: %s", args[1])
	}

	return ex.Exec(ctx, args[1:]...)
}

func (subex SubCommandExecutor) Complete(ctx context.Context, cmd string) ([]string, error) {
	args := strings.Split(cmd, " ")
	var completions []string
	switch len(args) {
	case 0: // complete all sub commands when there are no args
		for name := range subex {
			completions = append(completions, name)
		}
	case 1: // one arg means user is trying to complete the name of a sub command
		for name := range subex {
			if strings.HasPrefix(name, args[0]) {
				completions = append(completions, name)
			}
		}
	default: // more than one arg means we can pass completion responibility to a sub command
		for name, ex := range subex {
			if name == args[0] {
				if c, ok := ex.(Completer); ok {
					return c.Complete(ctx, strings.TrimLeft(strings.TrimPrefix(cmd, args[0]), " "))
				}
			}
		}
	}
	return completions, nil
}
