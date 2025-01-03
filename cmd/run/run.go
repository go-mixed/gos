package run

import (
	"fmt"
	"github.com/c4milo/unpackit"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
	"gopkg.in/go-mixed/gos.v1/mod"
	"os"
	"path/filepath"
	"strings"
)

type runCmdOptions struct {
	path        string
	debug       bool
	vendorPath  string
	importPaths map[string]string
	pluginPaths []string

	realPath  string
	isDir     bool
	isArchive bool
}

func RunCmd() *cobra.Command {
	var options = runCmdOptions{}

	runCmd := &cobra.Command{
		Use:   "run [OPTIONS] <PATH> -- <script argument>",
		Short: "Execute a Go+ script file, or a Golang project",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			options.path = args[0]
			code, err := igopRun(options, args[1:])
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
			if code != 0 {
				os.Exit(code)
			}
		},
	}

	runCmd.PersistentFlags().BoolVarP(&options.debug, "debug", "V", false, "Print debug information")
	runCmd.PersistentFlags().StringToStringVarP(&options.importPaths, "import", "I", map[string]string{}, "The package to be imported, -I NAME=PATH -I NAME2=PATH2")
	runCmd.PersistentFlags().StringVar(&options.vendorPath, "vendor", "", "The path of vendor, default: <PATH>/vendor")
	runCmd.PersistentFlags().StringArrayVarP(&options.pluginPaths, "plugin", "p", nil, "the golang plugin path (only for linux)")
	runCmd.MarkPersistentFlagDirname("vendor")
	return runCmd
}

func initialPath(options runCmdOptions) (runCmdOptions, error) {
	var err error
	var stat os.FileInfo
	if stat, err = os.Stat(options.path); err != nil {
		return options, err
	}

	options.isDir = stat.IsDir()
	// 获取绝对路径
	options.path, _ = filepath.Abs(options.path)

	// 是tar.gz压缩文件
	if !options.isDir && (strings.HasSuffix(options.path, ".tar.gz") ||
		strings.HasSuffix(options.path, ".tar.bzip2") ||
		strings.HasSuffix(options.path, ".tar.xz") ||
		strings.HasSuffix(options.path, ".zip") ||
		strings.HasSuffix(options.path, ".tar")) {
		var f *os.File
		if f, err = os.Open(options.path); err != nil {
			return options, err
		}
		defer f.Close()
		options.realPath = filepath.Join(filepath.Dir(options.path), "__"+filepath.Base(options.path)+"__")
		if err = unpackit.Unpack(f, options.realPath); err != nil {
			return options, err
		}
		options.isArchive = true
	} else {
		options.realPath = options.path
	}

	// 配置了vendor，检查是否有效
	if options.vendorPath != "" {
		// 转为绝对路径
		if options.vendorPath, err = filepath.Abs(options.vendorPath); err != nil {
			return options, fmt.Errorf("[Vendor]%w", err)
		}

		if stat, err = os.Stat(filepath.Join(options.vendorPath, "modules.txt")); err != nil || stat.IsDir() { // 不存在vendor/modules.txt
			return options, fmt.Errorf("[Vendor]%w", multierr.Append(err, fmt.Errorf("%s/modules.txt is not a regular file", options.vendorPath)))
		}
	}
	return options, nil
}

func igopRun(options runCmdOptions, args []string) (int, error) {
	var err error
	if options, err = initialPath(options); err != nil {
		return -11, err
	}

	// 删除解压的文件夹
	defer func() {
		if options.isArchive {
			os.RemoveAll(options.realPath)
		}
	}()

	ctx := mod.NewContext(options.realPath, options.debug)

	// 加载importPaths
	for k, v := range options.importPaths {
		if err = ctx.AddImport(k, v); err != nil {
			return -12, err
		}
		if options.debug {
			fmt.Printf("# imported package [%s]%s\n", k, v)
		}
	}
	// 加载plugins
	if err = ctx.LoadPlugins(options.pluginPaths); err != nil {
		return -13, err
	}
	// 加载vendor
	if options.isArchive || options.isDir || options.vendorPath != "" {
		if err = ctx.LoadVendor(options.vendorPath); err != nil {
			return -14, err
		}
	}

	return ctx.RunMain(args)
}
