package main

import (
	_ "github.com/goplus/igop/pkg"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use: "igop",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	addRunCmd(rootCmd)
	RunReplCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
