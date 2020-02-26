package bash

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	base "github.com/palsivertsen/go-subcommands"
)

// Completer implements the base.Command interface for easy bash completion
type Completer struct {
	RootCMD base.Command
}

// Exec completion using RootCMD
func (c *Completer) Exec(ctx context.Context, args ...string) error {
	compl, err := Complete(ctx, c.RootCMD)
	if err != nil {
		return err
	}

	for _, v := range compl {
		fmt.Println(v)
	}
	return nil
}

// Name returns "bash-completer"
func (c *Completer) Name() string {
	return "bash-completer"
}

func (c *Completer) SubCommands() []base.Command {
	return nil
}

func (c *Completer) Flags() []string {
	return nil
}

// Complete command for bash
func Complete(ctx context.Context, cmd base.Command) ([]string, error) {
	env, envErr := parseCompletionEnvironment()
	if envErr != nil {
		return nil, envErr
	}

	// trim after cursor
	line := env.CompLine[:env.CompPoint]

	// first word is this command
	words := strings.Split(line, " ")[1:]

	return complete(ctx, cmd, words)
}

// complete will recursivly try to complete for given command
func complete(ctx context.Context, cmd base.Command, words []string) ([]string, error) {
	// empty completion
	if len(words) == 0 {
		return nil, errors.New("can not handle empty completions yet")
	}

	subs := cmd.SubCommands()
	word := words[0]

	// not the last completion
	if len(words) > 1 {
		for _, s := range subs {
			if s.Name() == word {
				return complete(ctx, s, words[1:])
			}
		}
		return nil, fmt.Errorf("found no commands for words: %v", words)
	}

	// last completion
	compl := make([]string, 0, len(subs))

	for _, s := range subs {
		if strings.HasPrefix(s.Name(), word) {
			compl = append(compl, s.Name())
		}
	}

	for _, f := range cmd.Flags() {
		if strings.HasPrefix(f, word) {
			compl = append(compl, f)
		}
	}

	sort.Strings(compl)
	return compl, nil
}

// complEnv is a healper struct for the bash completion enviornment
type complEnv struct {
	CompLine  string
	CompPoint int
}

// parseCompletionEnvironment parses completion variables from environment
func parseCompletionEnvironment() (complEnv, error) {
	line := os.Getenv("COMP_LINE")

	pointRaw := os.Getenv("COMP_POINT")

	point, pointErr := strconv.Atoi(pointRaw)
	if pointErr != nil {
		return complEnv{}, pointErr
	}

	return complEnv{
		CompLine:  line,
		CompPoint: point,
	}, nil
}
