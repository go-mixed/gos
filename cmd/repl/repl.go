package repl

import (
	"fmt"
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/repl"
	"github.com/goplus/xgo/env"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
	"io"
	"strings"

	_ "gopkg.in/go-mixed/gos.v1/mod"
)

type replOptions struct {
	debug bool
}

func ReplCmd() *cobra.Command {
	var options replOptions

	replCmd := &cobra.Command{
		Use:   "repl",
		Short: "run a gos REPL " + env.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			gosRepl(options)
		},
	}

	replCmd.PersistentFlags().BoolVarP(&options.debug, "debug", "V", false, "print debug information")
	return replCmd
}

// LinerUI implements repl.UI interface.
type LinerUI struct {
	state  *liner.State
	prompt string
}

// SetPrompt is required by repl.UI interface.
func (u *LinerUI) SetPrompt(prompt string) {
	u.prompt = prompt
}

// Printf is required by repl.UI interface.
func (u *LinerUI) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func gosRepl(options replOptions) {
	fmt.Printf("Gos REPL %s\n", env.Version())
	state := liner.NewLiner()
	defer state.Close()

	// state.SetCtrlCAborts(true)
	state.SetMultiLineMode(true)
	state.SetCompleter(func(line string) []string {
		if strings.TrimSpace(line) == "" {
			return []string{line + "    "}
		}
		return nil
	})
	ui := &LinerUI{state: state}
	var mode ixgo.Mode
	if options.debug {
		mode |= ixgo.EnableDumpInstr | ixgo.EnableTracing
	}

	var r *repl.REPL
	ixgo.RegisterCustomBuiltin("exit", func() {
		r.Interp().Exit(0)
	})
	r = repl.NewREPL(mode)
	r.SetUI(ui)
	r.SetFileName("main.xgo") // support xgo

	for {
		line, err := ui.state.Prompt(ui.prompt)
		if err != nil {
			if err == liner.ErrPromptAborted || err == io.EOF {
				fmt.Printf("exit\n")
				break
			}
			fmt.Printf("Problem reading line: %v\n", err)
			continue
		}
		if line != "" {
			state.AppendHistory(line)
		}
		err = r.Run(line)
		switch e := err.(type) {
		case nil:
			//
		case ixgo.ExitError:
			fmt.Printf("exit %v\n", int(e))
			return
		default:
			fmt.Printf("error: %v\n", err)
		}
	}
}
