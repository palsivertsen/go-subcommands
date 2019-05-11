package bash

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/palsivertsen/go-subcommands"
)

func Complete(ctx context.Context, cmd subcommands.Command) ([]string, error) {
	env, envErr := getCompletionEnvironment()
	if envErr != nil {
		return nil, envErr
	}

	// trim after cursor
	line := env.CompLine[:env.CompPoint]

	// first word is this command
	words := strings.Split(line, " ")[1:]

	return complete(ctx, cmd, words)
}

func complete(ctx context.Context, cmd subcommands.Command, words []string) ([]string, error) {
	// empty completion
	if len(words) == 0 {
		return nil, errors.New("can not handle empty completions yet")
	}

	// flag completion
	if strings.HasPrefix(words[0], "-") {
		return nil, errors.New("can not handle flag completions yet")
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

	sort.Strings(compl)
	return compl, nil
}

type complEnv struct {
	CompLine  string
	CompPoint int
}

func getCompletionEnvironment() (complEnv, error) {
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
