// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package number

import (
	q "golang.org/x/text/number"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "number",
		Path: "golang.org/x/text/number",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"golang.org/x/text/feature/plural":  "plural",
			"golang.org/x/text/internal/format": "format",
			"golang.org/x/text/internal/number": "number",
			"golang.org/x/text/language":        "language",
			"strings":                           "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"FormatFunc": reflect.TypeOf((*q.FormatFunc)(nil)).Elem(),
			"Formatter":  reflect.TypeOf((*q.Formatter)(nil)).Elem(),
			"Option":     reflect.TypeOf((*q.Option)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decimal":           reflect.ValueOf(q.Decimal),
			"Engineering":       reflect.ValueOf(q.Engineering),
			"FormatWidth":       reflect.ValueOf(q.FormatWidth),
			"IncrementString":   reflect.ValueOf(q.IncrementString),
			"MaxFractionDigits": reflect.ValueOf(q.MaxFractionDigits),
			"MaxIntegerDigits":  reflect.ValueOf(q.MaxIntegerDigits),
			"MinFractionDigits": reflect.ValueOf(q.MinFractionDigits),
			"MinIntegerDigits":  reflect.ValueOf(q.MinIntegerDigits),
			"NewFormat":         reflect.ValueOf(q.NewFormat),
			"NoSeparator":       reflect.ValueOf(q.NoSeparator),
			"Pad":               reflect.ValueOf(q.Pad),
			"PatternOverrides":  reflect.ValueOf(q.PatternOverrides),
			"PerMille":          reflect.ValueOf(q.PerMille),
			"Percent":           reflect.ValueOf(q.Percent),
			"Precision":         reflect.ValueOf(q.Precision),
			"Scale":             reflect.ValueOf(q.Scale),
			"Scientific":        reflect.ValueOf(q.Scientific),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
