// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package simplifiedchinese

import (
	q "golang.org/x/text/encoding/simplifiedchinese"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "simplifiedchinese",
		Path: "golang.org/x/text/encoding/simplifiedchinese",
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
			"All":      reflect.ValueOf(&q.All),
			"GB18030":  reflect.ValueOf(&q.GB18030),
			"GBK":      reflect.ValueOf(&q.GBK),
			"HZGB2312": reflect.ValueOf(&q.HZGB2312),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
