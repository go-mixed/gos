// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package main

import (
	q "golang.org/x/text/collate/tools/colcmp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "main",
		Path: "golang.org/x/text/collate/tools/colcmp",
		Deps: map[string]string{
			"bytes":                          "bytes",
			"flag":                           "flag",
			"fmt":                            "fmt",
			"golang.org/x/text/collate":      "collate",
			"golang.org/x/text/language":     "language",
			"golang.org/x/text/unicode/norm": "norm",
			"io":                             "io",
			"log":                            "log",
			"math":                           "math",
			"math/rand":                      "rand",
			"os":                             "os",
			"runtime/pprof":                  "pprof",
			"sort":                           "sort",
			"strconv":                        "strconv",
			"strings":                        "strings",
			"text/template":                  "template",
			"time":                           "time",
			"unicode":                        "unicode",
			"unicode/utf16":                  "utf16",
			"unicode/utf8":                   "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Collator": reflect.TypeOf((*q.Collator)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CollatorFactory": reflect.TypeOf((*q.CollatorFactory)(nil)).Elem(),
			"Command":         reflect.TypeOf((*q.Command)(nil)).Elem(),
			"Context":         reflect.TypeOf((*q.Context)(nil)).Elem(),
			"Input":           reflect.TypeOf((*q.Input)(nil)).Elem(),
			"Test":            reflect.TypeOf((*q.Test)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AddFactory": reflect.ValueOf(q.AddFactory),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
