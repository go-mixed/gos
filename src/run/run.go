package run

import (
	"fmt"
	"github.com/c4milo/unpackit"
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
	"igop/src/mod"
	"os"
	"path/filepath"
	"strings"
)

type runOptions struct {
	path        string
	projectDir  string
	isDir       bool
	isArchive   bool
	debug       bool
	vendorPath  string
	importPaths map[string]string
	pluginPaths []string
}

func RunCmd() *cobra.Command {
	var options = runOptions{}

	runCmd := &cobra.Command{
		Use:   "run [OPTIONS] <PATH> -- <script argument>",
		Short: "Execute a Go+ script file, or a Golang project",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code, err := igoRun(args[0], options, args[1:])
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

func gopBuildDir(ctx *igop.Context, path string) error {
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666)
}

func buildOptions(path string, options *runOptions) error {
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
		options.projectDir = filepath.Join(filepath.Dir(options.path), "__"+filepath.Base(options.path)+"__")
		if err = unpackit.Unpack(f, options.projectDir); err != nil {
			return err
		}
		options.isArchive = true
	} else if options.isDir {
		options.projectDir = options.path
	} else {
		options.projectDir = filepath.Dir(options.path)
	}

	// 配置了vendor，检查是否有效
	if options.vendorPath != "" {
		// 压缩包模式，并且vendor非绝对路径，则将压缩包目录附加在前
		if options.isArchive && !filepath.IsAbs(options.vendorPath) {
			options.vendorPath = filepath.Join(options.projectDir, options.vendorPath)
		}

		// 变为绝对路径
		if options.vendorPath, err = filepath.Abs(options.vendorPath); err != nil {
			return fmt.Errorf("[Vendor]%w", err)
		}

		if stat, err = os.Stat(filepath.Join(options.vendorPath, "modules.txt")); err != nil || stat.IsDir() { // 不存在vendor/modules.txt
			return fmt.Errorf("[Vendor]%w", multierr.Append(err, fmt.Errorf("%s/modules.txt is not a regular file", options.vendorPath)))
		}
	}
	return nil
}

func igoRun(path string, options runOptions, args []string) (int, error) {
	// 删除解压的文件夹
	defer func() {
		if options.isArchive {
			os.RemoveAll(options.projectDir)
		}
	}()

	var err error
	var code int
	var mode = igop.EnablePrintAny
	if options.debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	// 解压、预读modules
	if err = buildOptions(path, &options); err != nil {
		return -1, err
	}

	ctx := igop.NewContext(mode)

	for k, v := range options.importPaths {
		if err = ctx.AddImport(k, v); err != nil {
			return -2, err
		}
		if options.debug {
			fmt.Printf("# imported package [%s]%s\n", k, v)
		}
	}

	modules := mod.NewModules(options.projectDir)
	modules.SetLookup(ctx)
	// 加载plugins
	if err = modules.LoadPlugins(options.pluginPaths); err != nil {
		return -3, err

	}
	// 加载vendor
	if options.vendorPath != "" {
		if err = modules.LoadVendor(options.vendorPath); err != nil {
			return -4, err
		}
	}

	if options.isDir || options.isArchive {
		// 读取go.mod
		if err = modules.LoadGoMod(mod.GetModPath(options.projectDir)); err != nil {
			return -5, err
		}

		// 检查目录下是否有gop文件
		if containsExt(options.projectDir, ".gop") {
			if containsSubModules(options.projectDir) {
				return -6, fmt.Errorf("*.gop is not allowed in project mode with 3rd party modules")
			}
			if err = gopBuildDir(ctx, options.projectDir); err != nil {
				return -4, err
			}
		}
		code, err = ctx.Run(options.projectDir, args)
	} else {
		_path := options.path
		ext := filepath.Ext(options.path)
		var buf []byte
		if buf, err = os.ReadFile(options.path); err != nil {
			return -7, err
		}

		// 修改后缀
		if ext != ".go" && ext != ".gop" {
			_path = _path + ".gop"
		}

		// 修改bang line为golang支持的注释, 参考 https://github.com/erning/gorun/blob/master/gorun.go#L167
		if len(buf) > 2 && buf[0] == '#' && buf[1] == '!' {
			buf[0] = '/'
			buf[1] = '/'
		}
		code, err = ctx.RunFile(_path, buf, args)
	}

	if err != nil {
		return code, fmt.Errorf("exit code %d: %w", code, err)
	}
	return 0, nil
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

func containsSubModules(projectPath string) bool {
	// go.mod
	if stat, err := os.Stat(filepath.Join(projectPath, "go.mod")); err == nil && !stat.IsDir() {
		return true
	}

	// 查找子目录包含 *.go，只做了简单的查询
	if f, err := os.Open(projectPath); err == nil {
		defer f.Close()
		fis, _ := f.Readdir(-1)
		for _, fi := range fis {
			if fi.IsDir() && containsExt(filepath.Join(projectPath, fi.Name()), ".go") {
				return true
			}
		}
	}

	return false
}
