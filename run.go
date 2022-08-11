package main

import (
	"bufio"
	"fmt"
	"github.com/c4milo/unpackit"
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
	isArchive   bool
	debug       bool
	vendorPath  string
	importPaths map[string]string
}

func addRunCmd(rootCmd *cobra.Command) {
	var runOptions = runOptions{importPaths: map[string]string{}}

	runCmd := &cobra.Command{
		Use:   "run [OPTIONS] [PATH] -- [GOP ARG...]",
		Short: "execute a go+ script file, or a folder of golang",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// 如果不传递 [PATH] 则将工作目录作为[PATH]
			if len(args) == 0 {
				p, _ := os.Getwd()
				args = append(args, p)
			}

			code, err := igoRun(args[0], runOptions, args[1:])
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
			if code != 0 {
				os.Exit(code)
			}
		},
	}

	runCmd.PersistentFlags().BoolVarP(&runOptions.debug, "debug", "V", false, "print debug information")

	runCmd.PersistentFlags().StringVar(&runOptions.vendorPath, "vendor", "", "path of vendor, default: [PATH]/vendor")
	runCmd.MarkPersistentFlagDirname("vendor")
	rootCmd.AddCommand(runCmd)
}

func gopBuildDir(ctx *igop.Context, path string) error {
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666)
}

func build(path string, options *runOptions) error {
	var err error
	var stat os.FileInfo

	if stat, err = os.Stat(path); err != nil {
		return err
	}

	options.isDir = stat.IsDir()
	// 获取绝对路径
	options.path, _ = filepath.Abs(path)

	// 是tar.gz压缩文件
	if !options.isDir && (strings.HasSuffix(options.path, ".tar.gz") ||
		strings.HasSuffix(options.path, ".tar.bzip2") ||
		strings.HasSuffix(options.path, ".tar.xz") ||
		strings.HasSuffix(options.path, ".zip") ||
		strings.HasSuffix(options.path, ".tar")) {
		f, err := os.Open(options.path)
		if err != nil {
			return err
		}
		defer f.Close()
		options.dir = filepath.Join(filepath.Dir(options.path), "__"+filepath.Base(options.path)+"__")
		if _, err = unpackit.Unpack(f, options.dir); err != nil {
			return err
		}
		options.isArchive = true
	} else if options.isDir {
		options.dir = options.path
	} else {
		options.dir = filepath.Dir(options.path)
	}

	// 查找项目中是否有vendor目录
	if options.vendorPath == "" {
		vp := filepath.Join(options.dir, "vendor")
		if stat, err = os.Stat(filepath.Join(vp, "modules.txt")); err == nil && !stat.IsDir() { // 项目中存在vendor/modules.txt
			options.vendorPath = vp
		}
	}

	if options.vendorPath != "" {
		// 压缩包模式，并且vendor非绝对路径，则将压缩包目录附加在前
		if options.isArchive && !filepath.IsAbs(options.vendorPath) {
			options.vendorPath = filepath.Join(options.dir, options.vendorPath)
		}
		modulesTxt := filepath.Join(options.vendorPath, "modules.txt")

		if stat, err = os.Stat(modulesTxt); err != nil || stat.IsDir() { // vendor/modules.txt不存在
			return fmt.Errorf("%s/modules.txt is not exist %w", options.vendorPath, err)
		}
		options.vendorPath, _ = filepath.Abs(options.vendorPath)

		f, err := os.Open(modulesTxt)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.Contains(line, "#") {
				options.importPaths[line] = filepath.Join(options.vendorPath, line)
			}
		}

		if err = scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func igoRun(path string, runOptions runOptions, args []string) (int, error) {
	// 删除解压的文件夹
	defer func() {
		if runOptions.isArchive {
			os.RemoveAll(runOptions.dir)
		}
	}()

	var err error
	var code int
	var mode = igop.EnablePrintAny
	if runOptions.debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	// 解压、预读modules
	if err = build(path, &runOptions); err != nil {
		return -1, err
	}

	ctx := igop.NewContext(mode)

	for k, v := range runOptions.importPaths {
		if err = ctx.AddImport(k, v); err != nil {
			return -1, err
		}
		if runOptions.debug {
			fmt.Printf("# imported package [%s]%s\n", k, v)
		}
	}

	if runOptions.isDir || runOptions.isArchive {
		gopCount := countByExt(runOptions.dir, ".gop")
		if gopCount == 1 {
			if err = gopBuildDir(ctx, runOptions.dir); err != nil {
				return -1, err
			}
		} else if gopCount > 1 {
			return -1, fmt.Errorf("there can be one *.gop in PROJECT compile mode")
		}
		code, err = ctx.Run(runOptions.dir, args)
	} else {
		//var buf []byte
		//if buf, err = os.ReadFile(runOptions.path); err != nil {
		//	return err
		//}
		code, err = ctx.RunFile(runOptions.path, nil, args)
	}

	if err != nil {
		return -1, fmt.Errorf("exit code %d: %w", code, err)
	}
	return 0, nil
}

func countByExt(srcDir string, ext string) int {
	extCount := 0
	if f, err := os.Open(srcDir); err == nil {
		defer f.Close()
		fis, _ := f.Readdir(-1)
		for _, fi := range fis {
			if !fi.IsDir() && filepath.Ext(fi.Name()) == ext {
				extCount++
			}
		}
	}
	return extCount
}
