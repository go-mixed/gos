// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package build

import (
	q "golang.org/x/text/collate/build"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "build",
		Path: "golang.org/x/text/collate/build",
		Deps: map[string]string{
			"fmt":                                "fmt",
			"golang.org/x/text/internal/colltab": "colltab",
			"golang.org/x/text/language":         "language",
			"golang.org/x/text/unicode/norm":     "norm",
			"hash/fnv":                           "fnv",
			"io":                                 "io",
			"log":                                "log",
			"reflect":                            "reflect",
			"sort":                               "sort",
			"strings":                            "strings",
			"unicode":                            "unicode",
			"unicode/utf8":                       "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Builder":   reflect.TypeOf((*q.Builder)(nil)).Elem(),
			"Tailoring": reflect.TypeOf((*q.Tailoring)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewBuilder": reflect.ValueOf(q.NewBuilder),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
