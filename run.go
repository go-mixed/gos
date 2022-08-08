package main

import (
	"bufio"
	"fmt"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

type runOptions struct {
	path        string
	dir         string
	isDir       bool
	debug       bool
	vendorPath  string
	importPaths map[string]string
}

func addRunCmd(rootCmd *cobra.Command) {
	var runOptions = runOptions{importPaths: map[string]string{}}

	runCmd := &cobra.Command{
		Use:   "run [OPTIONS] [PATH] -- [GOP ARG...]",
		Short: "execute a go+ script file, or a folder of golang",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var stat os.FileInfo
			path := args[0]

			if stat, err = os.Stat(path); err != nil {
				return err
			}

			runOptions.isDir = stat.IsDir()
			runOptions.path, _ = filepath.Abs(path)
			if stat.IsDir() {
				runOptions.dir = runOptions.path
			} else {
				runOptions.dir = filepath.Dir(runOptions.path)
			}

			if runOptions.vendorPath == "" {
				runOptions.vendorPath = filepath.Join(runOptions.dir, "vendor")
			}

			// vendor
			if stat, err = os.Stat(runOptions.vendorPath); err == nil {
				runOptions.vendorPath, _ = filepath.Abs(runOptions.vendorPath)
				if !stat.IsDir() {
					return nil
				}

				f, err := os.Open(filepath.Join(runOptions.vendorPath, "modules.txt"))
				if err != nil {
					return err
				}
				defer f.Close()
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					line := scanner.Text()
					if !strings.Contains(line, "#") {
						runOptions.importPaths[line] = filepath.Join(runOptions.vendorPath, line)
					}
				}

				if err = scanner.Err(); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return igoRun(runOptions, args[1:])
		},
	}

	runCmd.PersistentFlags().BoolVarP(&runOptions.debug, "debug", "V", false, "print debug information")
	runCmd.PersistentFlags().StringVar(&runOptions.vendorPath, "vendor", "", "path of vendor, default: [PATH]/vendor")

	rootCmd.AddCommand(runCmd)
}

func gopBuildDir(ctx *igop.Context, path string) error {
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666)
}

func igoRun(runOptions runOptions, args []string) error {
	var err error
	var code int
	var mode = igop.EnablePrintAny
	if runOptions.debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	ctx := igop.NewContext(mode)

	for k, v := range runOptions.importPaths {
		if err = ctx.AddImport(k, v); err != nil {
			return err
		}
		if runOptions.debug {
			fmt.Printf("# imported package [%s]%s\n", k, v)
		}
	}

	if runOptions.isDir {
		if containsExt(runOptions.dir, ".gop") {
			if err = gopBuildDir(ctx, runOptions.dir); err != nil {
				return err
			}
		}
		code, err = ctx.Run(runOptions.path, args)
	} else {
		//var buf []byte
		//if buf, err = os.ReadFile(runOptions.path); err != nil {
		//	return err
		//}
		code, err = ctx.RunFile(runOptions.path, nil, args)
	}

	if err != nil {
		return fmt.Errorf("exit code %d: %w", code, err)
	}
	return nil
}

func containsExt(srcDir string, ext string) bool {
	if f, err := os.Open(srcDir); err == nil {
		defer f.Close()
		fis, _ := f.Readdir(-1)
		for _, fi := range fis {
			if !fi.IsDir() && filepath.Ext(fi.Name()) == ext {
				return true
			}
		}
	}
	return false
}
