package main

import (
	_ "github.com/goplus/igop/pkg"
	"github.com/spf13/cobra"
	"igop/src/repl"
	"igop/src/run"
)

func main() {

	rootCmd := &cobra.Command{
		Use: "igop",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	run.AddRunCmd(rootCmd)
	repl.AddReplCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
