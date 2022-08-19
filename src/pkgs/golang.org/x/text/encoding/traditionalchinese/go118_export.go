// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package traditionalchinese

import (
	q "golang.org/x/text/encoding/traditionalchinese"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "traditionalchinese",
		Path: "golang.org/x/text/encoding/traditionalchinese",
		Deps: map[string]string{
			"golang.org/x/text/encoding":                     "encoding",
			"golang.org/x/text/encoding/internal":            "internal",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
			"unicode/utf8":                                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"All":  reflect.ValueOf(&q.All),
			"Big5": reflect.ValueOf(&q.Big5),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
