package main

import (
	"github.com/fly-studio/igop/src/exec"
	"github.com/fly-studio/igop/src/repl"
	"github.com/fly-studio/igop/src/run"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use: "igop",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(run.RunCmd())
	rootCmd.AddCommand(repl.ReplCmd())
	rootCmd.AddCommand(exec.ExecCmd())

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
