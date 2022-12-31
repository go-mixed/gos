package main

import (
	"github.com/spf13/cobra"
	"gopkg.in/go-mixed/igop.v1/cmd/exec"
	"gopkg.in/go-mixed/igop.v1/cmd/repl"
	"gopkg.in/go-mixed/igop.v1/cmd/run"
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
