package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	base "github.com/palsivertsen/go-subcommands"
	"github.com/palsivertsen/go-subcommands/bash"
)

/*
Simple example that uses a root command, an simple echo command and the bash completer.

To run the example you must compile a binary and set up the bash completer like so:

complete -C "<path to your binary> bash-completer" <your binary name>
*/

func main() {
	cmd := &rootCmd{}
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.Exec(ctx, flag.Args()...); err != nil {
		log.Fatal(err)
	}
}

type rootCmd struct {
}

func (c *rootCmd) Name() string {
	return "root"
}

func (c *rootCmd) Flags() []string {
	return nil
}

func (c *rootCmd) SubCommands() []base.Command {
	return []base.Command{
		&echo{},
		&bash.Completer{RootCMD: c},
	}
}

func (c *rootCmd) Exec(ctx context.Context, args ...string) error {
	if len(args) == 0 {
		fmt.Println("This is the root command")
		return nil
	}

	for _, s := range c.SubCommands() {
		if s.Name() == args[0] {
			return s.Exec(ctx, args[1:]...)
		}
	}
	return fmt.Errorf("command not found: %s", args[0])
}

// echo prints all args
type echo struct {
	echoStr string
}

func (c *echo) SubCommands() []base.Command {
	return nil
}

func (c *echo) Flags() []string {
	return []string{"-hello", "--world"}
}

func (c *echo) Name() string {
	return "echo"
}

func (c echo) Exec(ctx context.Context, args ...string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
