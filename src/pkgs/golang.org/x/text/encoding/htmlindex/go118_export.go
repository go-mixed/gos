// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package htmlindex

import (
	q "golang.org/x/text/encoding/htmlindex"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "htmlindex",
		Path: "golang.org/x/text/encoding/htmlindex",
		Deps: map[string]string{
			"errors":                                         "errors",
			"golang.org/x/text/encoding":                     "encoding",
			"golang.org/x/text/encoding/charmap":             "charmap",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/encoding/japanese":            "japanese",
			"golang.org/x/text/encoding/korean":              "korean",
			"golang.org/x/text/encoding/simplifiedchinese":   "simplifiedchinese",
			"golang.org/x/text/encoding/traditionalchinese":  "traditionalchinese",
			"golang.org/x/text/encoding/unicode":             "unicode",
			"golang.org/x/text/language":                     "language",
			"strings":                                        "strings",
			"sync":                                           "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Get":             reflect.ValueOf(q.Get),
			"LanguageDefault": reflect.ValueOf(q.LanguageDefault),
			"Name":            reflect.ValueOf(q.Name),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
