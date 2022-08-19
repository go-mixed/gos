// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package unicode

import (
	q "golang.org/x/text/encoding/unicode"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "unicode",
		Path: "golang.org/x/text/encoding/unicode",
		Deps: map[string]string{
			"bytes":                               "bytes",
			"errors":                              "errors",
			"golang.org/x/text/encoding":          "encoding",
			"golang.org/x/text/encoding/internal": "internal",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/internal/utf8internal":        "utf8internal",
			"golang.org/x/text/runes":                        "runes",
			"golang.org/x/text/transform":                    "transform",
			"unicode/utf16":                                  "utf16",
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
			"UTF8":          reflect.ValueOf(&q.UTF8),
			"UTF8BOM":       reflect.ValueOf(&q.UTF8BOM),
		},
		Funcs: map[string]reflect.Value{
			"BOMOverride": reflect.ValueOf(q.BOMOverride),
			"UTF16":       reflect.ValueOf(q.UTF16),
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
