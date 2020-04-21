package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/palsivertsen/go-subcommands"
)

func main() {
	echoExec := subcommands.ExecutorFunc(func(_ context.Context, args ...string) error {
		fmt.Println(strings.Join(args[1:], " "))
		return nil
	})

	greeterExec := subcommands.ExecutorFunc(func(_ context.Context, args ...string) error {
		if len(args) <= 1 {
			fmt.Println("Greetings!")
		} else {
			fmt.Printf("Greetings %s!", args[1])
		}
		return nil
	})

	rootExec := subcommands.SubCommandExecutor{
		"echo":  echoExec,
		"greet": greeterExec,
	}

	cli := subcommands.CLI{
		Executor: subcommands.WrapBash(rootExec),
	}

	if err := cli.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
