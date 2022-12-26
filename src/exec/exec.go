package exec

import (
	"bufio"
	"fmt"
	"github.com/goplus/igop"
	"github.com/spf13/cobra"
	"io"
	"os"

	_ "github.com/fly-studio/igop/src/mod"
)

type execOptions struct {
	debug       bool
	script      string
	scriptIsSet bool
}

func ExecCmd() *cobra.Command {
	var options = execOptions{}

	execCmd := &cobra.Command{
		Use:   "exec [-s | --script <code>] -- <script arguments>",
		Short: "execute script <code> from stdin, or argument of -s \"code\"",
		Run: func(cmd *cobra.Command, args []string) {
			options.scriptIsSet = cmd.PersistentFlags().Changed("script")

			code, err := igoExec(options, args)
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
			if code != 0 {
				os.Exit(code)
			}
		},
	}
	execCmd.PersistentFlags().StringVarP(&options.script, "script", "s", "", "exec the script")
	execCmd.PersistentFlags().BoolVarP(&options.debug, "debug", "V", false, "print debug information")
	return execCmd
}

func igoExec(options execOptions, args []string) (int, error) {
	var content []byte

	if options.scriptIsSet {
		content = []byte(options.script)
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil || (stat.Mode()&os.ModeCharDevice) != 0 {
			return -1, fmt.Errorf("must input a valid file or content, \"igop exec < 1.txt\"")
		}
		content, err = io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return -2, err
		}
	}

	var mode = igop.EnablePrintAny
	if options.debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	ctx := igop.NewContext(mode)

	return ctx.RunFile("main.gop", content, args)
}
