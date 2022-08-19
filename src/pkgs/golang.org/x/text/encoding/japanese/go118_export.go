// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package japanese

import (
	q "golang.org/x/text/encoding/japanese"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "japanese",
		Path: "golang.org/x/text/encoding/japanese",
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
			"All":       reflect.ValueOf(&q.All),
			"EUCJP":     reflect.ValueOf(&q.EUCJP),
			"ISO2022JP": reflect.ValueOf(&q.ISO2022JP),
			"ShiftJIS":  reflect.ValueOf(&q.ShiftJIS),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
