package cmd

import (
	"fmt"
	"github.com/c4milo/unpackit"
	"go.uber.org/multierr"
	"gopkg.in/go-mixed/gos.v1/mod"
	"os"
	"path/filepath"
	"strings"
)

func initialPath(options CmdOptions) (CmdOptions, error) {
	var err error
	var stat os.FileInfo
	if stat, err = os.Stat(options.Path); err != nil {
		return options, err
	}

	options.isDir = stat.IsDir()
	// 获取绝对路径
	options.Path, _ = filepath.Abs(options.Path)

	// 是tar.gz压缩文件
	if !options.isDir && (strings.HasSuffix(options.Path, ".tar.gz") ||
		strings.HasSuffix(options.Path, ".tar.bzip2") ||
		strings.HasSuffix(options.Path, ".tar.xz") ||
		strings.HasSuffix(options.Path, ".zip") ||
		strings.HasSuffix(options.Path, ".tar")) {
		var f *os.File
		if f, err = os.Open(options.Path); err != nil {
			return options, err
		}
		defer f.Close()
		options.realPath = filepath.Join(filepath.Dir(options.Path), "__"+filepath.Base(options.Path)+"__")
		if err = unpackit.Unpack(f, options.realPath); err != nil {
			return options, err
		}
		options.isArchive = true
	} else {
		options.realPath = options.Path
	}

	// 配置了vendor，检查是否有效
	if options.VendorPath != "" {
		// 转为绝对路径
		if options.VendorPath, err = filepath.Abs(options.VendorPath); err != nil {
			return options, fmt.Errorf("[Vendor]%w", err)
		}

		if stat, err = os.Stat(filepath.Join(options.VendorPath, "modules.txt")); err != nil || stat.IsDir() { // 不存在vendor/modules.txt
			return options, fmt.Errorf("[Vendor]%w", multierr.Append(err, fmt.Errorf("%s/modules.txt is not a regular file", options.VendorPath)))
		}
	}
	return options, nil
}

func IgopRun(options CmdOptions, args []string) (int, error) {
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

	ctx := mod.NewContext(options.realPath, options.Debug)

	// 加载importPaths
	for k, v := range options.ImportPaths {
		if err = ctx.AddImport(k, v); err != nil {
			return -12, err
		}
		if options.Debug {
			fmt.Printf("# imported package [%s]%s\n", k, v)
		}
	}
	// 加载plugins
	if err = ctx.LoadPlugins(options.PluginPaths); err != nil {
		return -13, err
	}
	// 加载vendor
	if options.isArchive || options.isDir || options.VendorPath != "" {
		if err = ctx.LoadVendor(options.VendorPath); err != nil {
			return -14, err
		}
	}

	return ctx.RunMain(args)
}
