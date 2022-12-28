package mod

import (
	"fmt"
	"github.com/goplus/igop"
	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/ssa"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

import (
	_ "github.com/fly-studio/igop/pkgs"
	_ "github.com/goplus/igop/pkg"
	_ "github.com/goplus/ipkg/github.com/modern-go/reflect2"
	_ "github.com/goplus/reflectx/icall/icall8192"

	_ "github.com/goplus/igop/gopbuild" // 注册gop后缀
)

type Module struct {
	Name      string
	Path      string
	Version   string
	GoVersion string
}

type Context struct {
	*igop.Context
	mainPackage *ssa.Package

	path    string
	modules map[string]*Module
	rKeys   []string
	debug   bool
}

func NewContext(projectPath string, debug bool) *Context {
	var mode = igop.EnablePrintAny
	if debug {
		mode |= igop.EnableTracing | igop.EnableDumpImports | igop.EnableDumpInstr
	}

	var ctx = &Context{
		Context: igop.NewContext(mode),
		modules: map[string]*Module{},
		path:    canonicalize(projectPath),
		debug:   debug,
	}

	ctx.Context.Lookup = ctx.lookup

	return ctx
}

func (ctx *Context) GetPath() string {
	return ctx.path
}

func (ctx *Context) IsDebug() bool {
	return ctx.debug
}

func (ctx *Context) GetIgop() *igop.Context {
	return ctx.Context
}

func (ctx *Context) GetMainPackage() *ssa.Package {
	return ctx.mainPackage
}

func (ctx *Context) GetModules() map[string]*Module {
	return ctx.modules
}

func (ctx *Context) RunMain(args []string) (int, error) {
	if ctx.mainPackage == nil {
		if err := ctx.Build(); err != nil {
			return -1, err
		}
	}

	if !isMainPkg(ctx.mainPackage) {
		return -3, fmt.Errorf("\"func main(){}\" is undefined in package \"%s\"", ctx.mainPackage.Pkg.Name())
	}

	return ctx.RunPkg(ctx.mainPackage, ctx.path, args)
}

func (ctx *Context) RunFunc(funcName string, args []igop.Value) (igop.Value, error) {
	if ctx.mainPackage == nil {
		if err := ctx.Build(); err != nil {
			return -1, err
		}
	}

	interp, err := ctx.NewInterp(ctx.mainPackage)
	if err != nil {
		return -2, err
	}
	return interp.RunFunc(funcName, args...)
}

func (ctx *Context) Build() error {
	var err error
	isDir := false
	if stat, err := os.Stat(ctx.path); err != nil {
		return err
	} else {
		isDir = stat.IsDir()
	}

	if isDir {
		// 读取go.mod
		if err = ctx.LoadGoMod(GetModPath(ctx.path)); err != nil {
			return err
		}

		// 检查目录下是否有gop文件
		if containsExt(ctx.path, ".gop") {
			if containsSubModules(ctx.path) {
				return errors.New("*.gop is not allowed in project mode with 3rd party modules")
			}
			if err = gopBuildDir(ctx.Context, ctx.path); err != nil {
				return err
			}
		}

		ctx.mainPackage, err = ctx.LoadDir(ctx.path, false)
	} else {
		_path := ctx.path
		ext := filepath.Ext(_path)
		var buf []byte
		if buf, err = os.ReadFile(_path); err != nil {
			return errors.WithStack(err)
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

		ctx.mainPackage, err = ctx.LoadFile(_path, buf)
	}

	return errors.WithStack(err)
}

func (ctx *Context) resortKeys() {
	for k := range ctx.modules {
		ctx.rKeys = append(ctx.rKeys, k)
	}
	// rKeys 倒序排序
	sort.Slice(ctx.rKeys, func(i, j int) bool {
		return ctx.rKeys[i] > ctx.rKeys[j]
	})
}

func (ctx *Context) LoadGoMod(goModPath string) error {
	goModPath = canonicalize(goModPath)
	// go.mod存在
	if stat, err := os.Stat(goModPath); err == nil && !stat.IsDir() {
		if err = ctx.parseGoMod(goModPath); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) LoadVendor(vendorPath string) error {
	if vendorPath == "" { // vendor 目录没有传递，尝试使用项目下的
		vendorPath = filepath.Join(ctx.path, "vendor")
	}

	vendorPath = canonicalize(vendorPath)

	// vendor/modules.txt文件存在
	if stat, err := os.Stat(filepath.Join(vendorPath, "modules.txt")); err == nil && !stat.IsDir() {
		if err = ctx.parseVendor(vendorPath); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) LoadModule(moduleName string, path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	ctx.modules[moduleName] = &Module{
		Name: moduleName,
		Path: path,
	}

	ctx.resortKeys()
	return nil
}

func (ctx *Context) parseGoMod(goModPath string) error {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return errors.WithStack(err)
	}
	f, err := modfile.Parse(goModPath, data, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	if f.Module == nil {
		// No module declaration. Must add module path.
		return errors.New("no module declaration in go.mod. To specify the module path:\n\tgo mod edit -module=example.com/mod")
	}

	goVersion, err := modFileGoVersion(f)
	if err != nil {
		return err
	}

	ctx.modules[f.Module.Mod.Path] = &Module{
		Name:      f.Module.Mod.Path,
		Path:      ctx.path,
		Version:   f.Module.Mod.Version,
		GoVersion: goVersion,
	}

	ctx.resortKeys()

	return nil
}

func (ctx *Context) parseVendor(vendorPath string) error {
	vendorList, err := readVendorList(vendorPath)
	if err != nil {
		return errors.Wrapf(err, "[Vendor]")
	}

	for k, v := range vendorList.vendorMeta {
		ctx.modules[k.Path] = &Module{
			Name:      k.Path,
			Path:      filepath.Join(vendorPath, k.Path),
			Version:   k.Version,
			GoVersion: v.GoVersion,
		}
	}

	ctx.resortKeys()

	return nil
}

func (ctx *Context) lookup(root, pkg string) (dir string, found bool) {

	module, ok := ctx.modules[pkg]
	if ok {
		return module.Path, ok
	}

	// 因为是倒序排列，故第一个匹配项是最长匹配
	for _, v := range ctx.rKeys {
		if strings.HasPrefix(pkg, v+"/") {
			module = ctx.modules[v]
			break
		}
	}

	if module != nil && module.Path != "" {
		return filepath.Join(module.Path, pkg[len(module.Name+"/"):]), true
	}

	return "", false
}
