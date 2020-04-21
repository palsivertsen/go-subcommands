package subcommands

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WrapBash(e Executor) Executor {
	return ExecutorFunc(func(ctx context.Context, args ...string) error {
		c, ok := e.(Completer)
		if !ok {
			return nil
		}
		if len(args) <= 1 {
			return e.Exec(ctx)
		}
		if args[1] != "bash-completion" {
			return e.Exec(ctx, args...)
		}

		cmdIEnv := os.Getenv("COMP_POINT")
		cmdI, err := strconv.Atoi(cmdIEnv)
		if err != nil {
			return fmt.Errorf(`bash env COMP_POINT not set to a valid int: "%s"`, cmdIEnv)
		}

		cmd := os.Getenv("COMP_LINE")
		if len(cmd) < cmdI {
			return fmt.Errorf(`bash env COMP_LINE lenght (%d) is shorter than index specified in COMP_POINT (%d)`, len(cmd), cmdI)
		}
		baseCMDName := strings.Split(cmd, " ")[0]
		s, err := c.Complete(ctx, string([]rune(cmd)[len(baseCMDName)+1:cmdI]))
		if err != nil {
			return err
		}
		for _, v := range s {
			fmt.Println(v)
		}
		return nil
	})
}
