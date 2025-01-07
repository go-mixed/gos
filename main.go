package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/go-mixed/gos.v1/cmd"
	"gopkg.in/go-mixed/gos.v1/cmd/repl"
	"os"
)

func main() {

	var options cmd.CmdOptions
	rootCmd := &cobra.Command{
		Use: "gos [OPTIONS] <PATH?> [-s | --script <code>] -- <script argument>",
		Long: `
Run a Go+ script file, or run a Golang project if <PATH> set.
Execute script if <code> from stdin, or -s "code".
`,
		Args: cobra.ArbitraryArgs, // 允许任意参数，必须设置这个，不然无法添加子命令
		Run: func(_cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				options.Path = args[0]
				code, err := cmd.IgopRun(options, args[1:])
				if err != nil {
					fmt.Fprint(os.Stderr, err.Error())
				}
				if code != 0 {
					os.Exit(code)
				}
			} else {
				options.ScriptIsSet = _cmd.Flags().Changed("script")

				code, err := cmd.IgopExec(options, args)
				if err != nil {
					fmt.Fprint(os.Stderr, err.Error())
				}
				if code != 0 {
					os.Exit(code)
				}
			}
		},
	}

	rootCmd.AddCommand(repl.ReplCmd())

	rootCmd.Flags().BoolVarP(&options.Debug, "debug", "V", false, "Print debug information")
	rootCmd.Flags().StringToStringVarP(&options.ImportPaths, "import", "I", map[string]string{}, "The package to be imported, -I NAME=PATH -I NAME2=PATH2")
	rootCmd.Flags().StringVar(&options.VendorPath, "vendor", "", "The path of vendor, default: <PATH>/vendor")
	rootCmd.Flags().StringArrayVarP(&options.PluginPaths, "plugin", "p", nil, "the golang plugin path (only for linux)")
	rootCmd.Flags().StringVarP(&options.Script, "script", "s", "", "exec the script")

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
