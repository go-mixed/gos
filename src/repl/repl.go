package repl

import (
	"fmt"
	"github.com/goplus/gop/env"
	"github.com/goplus/igop"
	"github.com/goplus/igop/repl"
	"github.com/peterh/liner"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

type replOptions struct {
	debug bool
}

func AddReplCmd(rootCmd *cobra.Command) {
	var replOptions replOptions

	replCmd := &cobra.Command{
		Use:   "repl",
		Short: "run a go+ REPL " + env.Version(),
		Run: func(cmd *cobra.Command, args []string) {
			runRepl(replOptions)
		},
	}

	replCmd.PersistentFlags().BoolVarP(&replOptions.debug, "debug", "V", false, "print debug information")

	rootCmd.AddCommand(replCmd)
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

func runRepl(replOptions replOptions) {
	fmt.Printf("iGo+ REPL %s\n", env.Version())
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
	var mode igop.Mode
	if replOptions.debug {
		mode |= igop.EnableDumpInstr | igop.EnableTracing
	}

	var r *repl.REPL
	igop.RegisterCustomBuiltin("exit", func() {
		r.Interp().Exit(0)
	})
	r = repl.NewREPL(mode)
	r.SetUI(ui)
	r.SetFileName("main.gop") // support go+

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
		case igop.ExitError:
			fmt.Printf("exit %v\n", int(e))
			return
		default:
			fmt.Printf("error: %v\n", err)
		}
	}
}
