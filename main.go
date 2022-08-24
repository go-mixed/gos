package main

import (
	"github.com/spf13/cobra"
	"igop/src/exec"
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

	rootCmd.AddCommand(run.RunCmd())
	rootCmd.AddCommand(repl.ReplCmd())
	rootCmd.AddCommand(exec.ExecCmd())

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
