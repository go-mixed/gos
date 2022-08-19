// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package charmap

import (
	q "golang.org/x/text/encoding/charmap"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "charmap",
		Path: "golang.org/x/text/encoding/charmap",
		Deps: map[string]string{
			"golang.org/x/text/encoding":                     "encoding",
			"golang.org/x/text/encoding/internal":            "internal",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
			"unicode/utf8":                                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Charmap": reflect.TypeOf((*q.Charmap)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"All":               reflect.ValueOf(&q.All),
			"CodePage037":       reflect.ValueOf(&q.CodePage037),
			"CodePage1047":      reflect.ValueOf(&q.CodePage1047),
			"CodePage1140":      reflect.ValueOf(&q.CodePage1140),
			"CodePage437":       reflect.ValueOf(&q.CodePage437),
			"CodePage850":       reflect.ValueOf(&q.CodePage850),
			"CodePage852":       reflect.ValueOf(&q.CodePage852),
			"CodePage855":       reflect.ValueOf(&q.CodePage855),
			"CodePage858":       reflect.ValueOf(&q.CodePage858),
			"CodePage860":       reflect.ValueOf(&q.CodePage860),
			"CodePage862":       reflect.ValueOf(&q.CodePage862),
			"CodePage863":       reflect.ValueOf(&q.CodePage863),
			"CodePage865":       reflect.ValueOf(&q.CodePage865),
			"CodePage866":       reflect.ValueOf(&q.CodePage866),
			"ISO8859_1":         reflect.ValueOf(&q.ISO8859_1),
			"ISO8859_10":        reflect.ValueOf(&q.ISO8859_10),
			"ISO8859_13":        reflect.ValueOf(&q.ISO8859_13),
			"ISO8859_14":        reflect.ValueOf(&q.ISO8859_14),
			"ISO8859_15":        reflect.ValueOf(&q.ISO8859_15),
			"ISO8859_16":        reflect.ValueOf(&q.ISO8859_16),
			"ISO8859_2":         reflect.ValueOf(&q.ISO8859_2),
			"ISO8859_3":         reflect.ValueOf(&q.ISO8859_3),
			"ISO8859_4":         reflect.ValueOf(&q.ISO8859_4),
			"ISO8859_5":         reflect.ValueOf(&q.ISO8859_5),
			"ISO8859_6":         reflect.ValueOf(&q.ISO8859_6),
			"ISO8859_6E":        reflect.ValueOf(&q.ISO8859_6E),
			"ISO8859_6I":        reflect.ValueOf(&q.ISO8859_6I),
			"ISO8859_7":         reflect.ValueOf(&q.ISO8859_7),
			"ISO8859_8":         reflect.ValueOf(&q.ISO8859_8),
			"ISO8859_8E":        reflect.ValueOf(&q.ISO8859_8E),
			"ISO8859_8I":        reflect.ValueOf(&q.ISO8859_8I),
			"ISO8859_9":         reflect.ValueOf(&q.ISO8859_9),
			"KOI8R":             reflect.ValueOf(&q.KOI8R),
			"KOI8U":             reflect.ValueOf(&q.KOI8U),
			"Macintosh":         reflect.ValueOf(&q.Macintosh),
			"MacintoshCyrillic": reflect.ValueOf(&q.MacintoshCyrillic),
			"Windows1250":       reflect.ValueOf(&q.Windows1250),
			"Windows1251":       reflect.ValueOf(&q.Windows1251),
			"Windows1252":       reflect.ValueOf(&q.Windows1252),
			"Windows1253":       reflect.ValueOf(&q.Windows1253),
			"Windows1254":       reflect.ValueOf(&q.Windows1254),
			"Windows1255":       reflect.ValueOf(&q.Windows1255),
			"Windows1256":       reflect.ValueOf(&q.Windows1256),
			"Windows1257":       reflect.ValueOf(&q.Windows1257),
			"Windows1258":       reflect.ValueOf(&q.Windows1258),
			"Windows874":        reflect.ValueOf(&q.Windows874),
			"XUserDefined":      reflect.ValueOf(&q.XUserDefined),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
