package base

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagCompleter_Complete(t *testing.T) {
	tests := map[string]struct {
		fs *flag.FlagSet
		r  []string
	}{
		"root -s": {
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("", flag.ExitOnError)
				fs.String("s1", "", "")
				fs.String("s2", "", "")
				return fs
			}(),
			r: []string{"s1", "s2"},
		},
		"root -": {
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("", flag.ExitOnError)
				fs.String("s1", "", "")
				fs.String("s2", "", "")
				fs.String("r1", "", "")
				fs.String("r2", "", "")
				return fs
			}(),
			r: []string{"r1", "r2", "s1", "s2"},
		},
		"root -s ": {
			fs: func() *flag.FlagSet {
				fs := flag.NewFlagSet("", flag.ExitOnError)
				fs.String("s1", "", "")
				fs.String("s2", "", "")
				fs.String("r1", "", "")
				fs.String("r2", "", "")
				return fs
			}(),
			r: []string{"r1", "r2", "s1", "s2"},
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			unit := FlagCompleter{
				FS: tt.fs,
			}

			r := unit.Complete(name)
			assert.Equal(t, tt.r, r)
		})
	}
}

type strComp []string

func (s strComp) Complete(string) []string {
	return []string(s)
}

func TestSubcommandCompleter_Complete(t *testing.T) {
	tests := map[string]struct {
		subs map[string]*Command
		r    []string
	}{
		"root one": {
			subs: map[string]*Command{
				"one": &Command{
					Completer: strComp([]string{"ONE"}),
				},
				"two": &Command{
					Completer: strComp([]string{"TWO"}),
				},
			},
			r: []string{"one"},
		},
		"root one ": {
			subs: map[string]*Command{
				"one": &Command{
					Completer: strComp([]string{"ONE"}),
				},
				"two": &Command{
					Completer: strComp([]string{"TWO"}),
				},
			},
			r: []string{"ONE"},
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			unit := SubCommandCompleter{
				Subs: tt.subs,
			}

			r := unit.Complete(name)
			assert.Equal(t, tt.r, r)
		})
	}
}

func TestPrioritycompleter_Complete(t *testing.T) {
	tests := map[string]struct {
		cs []Completer
		r  []string
	}{
		"first": {
			cs: []Completer{
				strComp([]string{"first"}),
				strComp([]string{"last"}),
			},
			r: []string{"first"},
		},
		"last": {
			cs: []Completer{
				strComp{},
				strComp([]string{"last"}),
			},
			r: []string{"last"},
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			unit := PriorityCompleter{
				Completers: tt.cs,
			}

			r := unit.Complete(name)
			assert.Equal(t, tt.r, r)
		})
	}
}
