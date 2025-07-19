package cmd

import (
	"bufio"
	"fmt"
	"github.com/goplus/ixgo"
	"io"
	"os"
)

func GosExec(options CmdOptions, args []string) (int, error) {
	var content []byte

	if options.ScriptIsSet {
		content = []byte(options.Script)
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil || (stat.Mode()&os.ModeCharDevice) != 0 {
			return -1, fmt.Errorf("must input a valid file or content, \"gos exec < 1.txt\"")
		}
		content, err = io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return -2, err
		}
	}

	var mode = ixgo.EnablePrintAny
	if options.Debug {
		mode |= ixgo.EnableTracing | ixgo.EnableDumpImports | ixgo.EnableDumpInstr
	}

	ctx := ixgo.NewContext(mode)

	return ctx.RunFile("main.xgo", content, args)
}
