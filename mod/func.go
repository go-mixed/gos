package mod

import (
	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/ssa"
	"os"
	"path/filepath"
)

func canonicalize(path string) string {
	if path == "" {
		return path
	}
	nPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	nPath = filepath.Clean(nPath)
	return nPath
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

func gopBuildDir(ctx *igop.Context, path string) error {
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(os.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666))
}

func isMainPkg(pkg *ssa.Package) bool {
	return pkg.Pkg.Name() == "main" && pkg.Func("main") != nil
}
