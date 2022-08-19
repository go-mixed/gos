// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ianaindex

import (
	q "golang.org/x/text/encoding/ianaindex"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ianaindex",
		Path: "golang.org/x/text/encoding/ianaindex",
		Deps: map[string]string{
			"errors":                                         "errors",
			"golang.org/x/text/encoding":                     "encoding",
			"golang.org/x/text/encoding/charmap":             "charmap",
			"golang.org/x/text/encoding/internal":            "internal",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/encoding/japanese":            "japanese",
			"golang.org/x/text/encoding/korean":              "korean",
			"golang.org/x/text/encoding/simplifiedchinese":   "simplifiedchinese",
			"golang.org/x/text/encoding/traditionalchinese":  "traditionalchinese",
			"golang.org/x/text/encoding/unicode":             "unicode",
			"golang.org/x/text/transform":                    "transform",
			"sort":                                           "sort",
			"strings":                                        "strings",
			"unicode":                                        "unicode",
			"unicode/utf8":                                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Index": reflect.TypeOf((*q.Index)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"IANA": reflect.ValueOf(&q.IANA),
			"MIB":  reflect.ValueOf(&q.MIB),
			"MIME": reflect.ValueOf(&q.MIME),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
