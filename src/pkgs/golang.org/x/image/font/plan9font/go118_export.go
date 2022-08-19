// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package plan9font

import (
	q "golang.org/x/image/font/plan9font"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "plan9font",
		Path: "golang.org/x/image/font/plan9font",
		Deps: map[string]string{
			"bytes":                         "bytes",
			"errors":                        "errors",
			"fmt":                           "fmt",
			"golang.org/x/image/font":       "font",
			"golang.org/x/image/math/fixed": "fixed",
			"image":                         "image",
			"image/color":                   "color",
			"log":                           "log",
			"strconv":                       "strconv",
			"strings":                       "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ParseFont":    reflect.ValueOf(q.ParseFont),
			"ParseSubfont": reflect.ValueOf(q.ParseSubfont),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
