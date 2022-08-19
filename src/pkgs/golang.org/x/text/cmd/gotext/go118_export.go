// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package main

import (
	q "golang.org/x/text/cmd/gotext"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "main",
		Path: "golang.org/x/text/cmd/gotext",
		Deps: map[string]string{
			"bufio":                              "bufio",
			"bytes":                              "bytes",
			"flag":                               "flag",
			"fmt":                                "fmt",
			"go/build":                           "build",
			"go/format":                          "format",
			"go/parser":                          "parser",
			"golang.org/x/text/language":         "language",
			"golang.org/x/text/message/pipeline": "pipeline",
			"golang.org/x/tools/go/buildutil":    "buildutil",
			"golang.org/x/tools/go/loader":       "loader",
			"io":                                 "io",
			"io/ioutil":                          "ioutil",
			"log":                                "log",
			"os":                                 "os",
			"strings":                            "strings",
			"sync":                               "sync",
			"text/template":                      "template",
			"unicode":                            "unicode",
			"unicode/utf8":                       "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Command": reflect.TypeOf((*q.Command)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
