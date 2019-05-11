package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	sub "github.com/palsivertsen/go-subcommands"
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
	sub.UnimplementedCommand
}

func (c *rootCmd) SubCommands() []sub.Command {
	return []sub.Command{
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
	sub.UnimplementedCommand
	echoStr string
}

func (c *echo) Name() string {
	return "echo"
}

func (c echo) Exec(ctx context.Context, args ...string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
