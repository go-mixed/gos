// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package encoding

import (
	q "golang.org/x/text/encoding"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "encoding",
		Path: "golang.org/x/text/encoding",
		Deps: map[string]string{
			"errors": "errors",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
			"io":                                             "io",
			"strconv":                                        "strconv",
			"unicode/utf8":                                   "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Encoding": reflect.TypeOf((*q.Encoding)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Decoder": reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"Encoder": reflect.TypeOf((*q.Encoder)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalidUTF8": reflect.ValueOf(&q.ErrInvalidUTF8),
			"Nop":            reflect.ValueOf(&q.Nop),
			"Replacement":    reflect.ValueOf(&q.Replacement),
			"UTF8Validator":  reflect.ValueOf(&q.UTF8Validator),
		},
		Funcs: map[string]reflect.Value{
			"HTMLEscapeUnsupported": reflect.ValueOf(q.HTMLEscapeUnsupported),
			"ReplaceUnsupported":    reflect.ValueOf(q.ReplaceUnsupported),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"ASCIISub": {"untyped rune", constant.MakeInt64(int64(q.ASCIISub))},
		},
	})
}
