package cmd

import (
	"bufio"
	"fmt"
	"github.com/goplus/igop"
	"io"
	"os"
)

func IgopExec(options CmdOptions, args []string) (int, error) {
	var content []byte

	if options.ScriptIsSet {
		content = []byte(options.Script)
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
	if options.Debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	ctx := igop.NewContext(mode)

	return ctx.RunFile("main.gop", content, args)
}
