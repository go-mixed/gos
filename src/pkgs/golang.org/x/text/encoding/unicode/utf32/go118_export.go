// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package utf32

import (
	q "golang.org/x/text/encoding/unicode/utf32"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "utf32",
		Path: "golang.org/x/text/encoding/unicode/utf32",
		Deps: map[string]string{
			"errors":                     "errors",
			"golang.org/x/text/encoding": "encoding",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
			"unicode/utf8":                                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"BOMPolicy":  reflect.TypeOf((*q.BOMPolicy)(nil)).Elem(),
			"Endianness": reflect.TypeOf((*q.Endianness)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"All":           reflect.ValueOf(&q.All),
			"ErrMissingBOM": reflect.ValueOf(&q.ErrMissingBOM),
		},
		Funcs: map[string]reflect.Value{
			"UTF32": reflect.ValueOf(q.UTF32),
		},
		TypedConsts: map[string]igop.TypedConst{
			"BigEndian":    {reflect.TypeOf(q.BigEndian), constant.MakeBool(bool(q.BigEndian))},
			"ExpectBOM":    {reflect.TypeOf(q.ExpectBOM), constant.MakeInt64(int64(q.ExpectBOM))},
			"IgnoreBOM":    {reflect.TypeOf(q.IgnoreBOM), constant.MakeInt64(int64(q.IgnoreBOM))},
			"LittleEndian": {reflect.TypeOf(q.LittleEndian), constant.MakeBool(bool(q.LittleEndian))},
			"UseBOM":       {reflect.TypeOf(q.UseBOM), constant.MakeInt64(int64(q.UseBOM))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
