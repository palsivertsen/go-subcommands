package subcommands

import (
	"context"
	"os"
)

// Executor represents a stand alone command
type Executor interface {
	Exec(ctx context.Context, args ...string) error
}

type ExecutorFunc func(ctx context.Context, args ...string) error

func (f ExecutorFunc) Exec(ctx context.Context, args ...string) error {
	return f(ctx, args...)
}

type CLI struct {
	Executor Executor
}

func (cli *CLI) Run(ctx context.Context) error {
	return cli.Executor.Exec(ctx, os.Args...)
}
