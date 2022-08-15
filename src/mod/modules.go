package mod

import (
	"fmt"
	"go.uber.org/multierr"
	"go/build"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Module struct {
	Name      string
	Path      string
	Version   string
	GoVersion string
}

type Modules struct {
	projectPath string
	vendorPath  string
	modules     map[string]*Module
	rkeys       []string
}

func canonicalize(path string) (string, error) {
	if path == "" {
		return path, nil
	}
	nPath, err := filepath.Abs(path)
	if err != nil {
		return path, err
	}
	nPath = filepath.Clean(nPath)
	return nPath, nil
}

func NewModules(projectPath string, vendorPath string) (*Modules, error) {
	var err error
	var m = &Modules{modules: map[string]*Module{}}

	if m.projectPath, err = canonicalize(projectPath); err != nil {
		return nil, err
	}

	if m.vendorPath, err = canonicalize(vendorPath); err != nil {
		return nil, err
	}

	if err = m.parseGoMod(m.projectPath); err != nil {
		return nil, err
	}
	if err = m.parseVendor(m.vendorPath); err != nil {
		return nil, err
	}

	for k, _ := range m.modules {
		m.rkeys = append(m.rkeys, k)
	}
	// rkeys 倒序排序
	sort.Slice(m.rkeys, func(i, j int) bool {
		return m.rkeys[i] > m.rkeys[j]
	})

	return m, nil
}

func (m *Modules) parseGoMod(projectPath string) error {
	var err error

	gomod := filepath.Join(projectPath, "go.mod")

	data, err := os.ReadFile(gomod)
	if err != nil {
		return err
	}
	f, err := modfile.Parse(gomod, data, nil)
	if err != nil {
		return err
	}
	if f.Module == nil {
		// No module declaration. Must add module path.
		return fmt.Errorf("no module declaration in go.mod. To specify the module path:\n\tgo mod edit -module=example.com/mod")
	}

	goVersion, err := modFileGoVersion(f)
	if err != nil {
		return err
	}

	m.modules[f.Module.Mod.Path] = &Module{
		Name:      f.Module.Mod.Path,
		Path:      projectPath,
		Version:   f.Module.Mod.Version,
		GoVersion: goVersion,
	}

	//for _, require := range f.Require {
	//	m.modules[require.Mod.Path] = &Module{
	//		Name:      require.Mod.Path,
	//		Path:      "",
	//		Version:   require.Mod.Version,
	//		GoVersion: "",
	//	}
	//}

	return nil
}

func (m *Modules) parseVendor(vendorPath string) error {
	if vendorPath == "" {
		return nil
	} else if stat, err := os.Stat(vendorPath); err != nil || !stat.IsDir() {
		return multierr.Append(err, fmt.Errorf("\"%s\" is not a valid directory", vendorPath))
	}

	vendorList, err := readVendorList(vendorPath)
	if err != nil {
		return err
	}

	for k, v := range vendorList.vendorMeta {
		m.modules[k.Path] = &Module{
			Name:      k.Path,
			Path:      filepath.Join(vendorPath, k.Path),
			Version:   k.Version,
			GoVersion: v.GoVersion,
		}
	}

	return nil
}

func (m *Modules) Lookup(root, pkg string) (dir string, found bool) {

	module, ok := m.modules[pkg]
	if ok {
		return module.Path, ok
	}

	// 因为是倒序排列，故第一个匹配项是最长匹配
	for _, v := range m.rkeys {
		if strings.HasPrefix(pkg, v+"/") {
			module = m.modules[v]
			break
		}
	}

	if module != nil && module.Path != "" {
		return filepath.Join(module.Path, pkg[len(module.Name+"/"):]), true
	}

	// 使用
	bp, err := build.Import(pkg, m.projectPath, build.FindOnly)
	if err == nil && bp.ImportPath == m.projectPath {
		return bp.Dir, true
	}

	panic(fmt.Errorf("package %s not found", pkg))
	return "", false
}
