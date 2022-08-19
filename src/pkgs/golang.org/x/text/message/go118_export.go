// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package message

import (
	q "golang.org/x/text/message"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "message",
		Path: "golang.org/x/text/message",
		Deps: map[string]string{
			"bytes":                             "bytes",
			"fmt":                               "fmt",
			"golang.org/x/text/feature/plural":  "plural",
			"golang.org/x/text/internal/format": "format",
			"golang.org/x/text/internal/number": "number",
			"golang.org/x/text/language":        "language",
			"golang.org/x/text/message/catalog": "catalog",
			"io":                                "io",
			"math":                              "math",
			"os":                                "os",
			"reflect":                           "reflect",
			"strconv":                           "strconv",
			"sync":                              "sync",
			"unicode/utf8":                      "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Reference": reflect.TypeOf((*q.Reference)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Option":  reflect.TypeOf((*q.Option)(nil)).Elem(),
			"Printer": reflect.TypeOf((*q.Printer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DefaultCatalog": reflect.ValueOf(&q.DefaultCatalog),
		},
		Funcs: map[string]reflect.Value{
			"Catalog":       reflect.ValueOf(q.Catalog),
			"Key":           reflect.ValueOf(q.Key),
			"MatchLanguage": reflect.ValueOf(q.MatchLanguage),
			"NewPrinter":    reflect.ValueOf(q.NewPrinter),
			"Set":           reflect.ValueOf(q.Set),
			"SetString":     reflect.ValueOf(q.SetString),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
