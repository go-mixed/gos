// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package charset

import (
	q "golang.org/x/net/html/charset"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "charset",
		Path: "golang.org/x/net/html/charset",
		Deps: map[string]string{
			"bytes":                                "bytes",
			"fmt":                                  "fmt",
			"golang.org/x/net/html":                "html",
			"golang.org/x/text/encoding":           "encoding",
			"golang.org/x/text/encoding/charmap":   "charmap",
			"golang.org/x/text/encoding/htmlindex": "htmlindex",
			"golang.org/x/text/transform":          "transform",
			"io":                                   "io",
			"mime":                                 "mime",
			"strings":                              "strings",
			"unicode/utf8":                         "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"DetermineEncoding": reflect.ValueOf(q.DetermineEncoding),
			"Lookup":            reflect.ValueOf(q.Lookup),
			"NewReader":         reflect.ValueOf(q.NewReader),
			"NewReaderLabel":    reflect.ValueOf(q.NewReaderLabel),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
