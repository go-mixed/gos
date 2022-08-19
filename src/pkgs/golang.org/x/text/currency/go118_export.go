// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package currency

import (
	q "golang.org/x/text/currency"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "currency",
		Path: "golang.org/x/text/currency",
		Deps: map[string]string{
			"errors":                            "errors",
			"fmt":                               "fmt",
			"golang.org/x/text/internal/format": "format",
			"golang.org/x/text/internal/language/compact": "compact",
			"golang.org/x/text/internal/tag":              "tag",
			"golang.org/x/text/language":                  "language",
			"io":                                          "io",
			"sort":                                        "sort",
			"time":                                        "time",
		},
		Interfaces: map[string]reflect.Type{
			"QueryIter": reflect.TypeOf((*q.QueryIter)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Amount":      reflect.TypeOf((*q.Amount)(nil)).Elem(),
			"Formatter":   reflect.TypeOf((*q.Formatter)(nil)).Elem(),
			"Kind":        reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"QueryOption": reflect.TypeOf((*q.QueryOption)(nil)).Elem(),
			"Unit":        reflect.TypeOf((*q.Unit)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"AUD":          reflect.ValueOf(&q.AUD),
			"Accounting":   reflect.ValueOf(&q.Accounting),
			"BRL":          reflect.ValueOf(&q.BRL),
			"CAD":          reflect.ValueOf(&q.CAD),
			"CHF":          reflect.ValueOf(&q.CHF),
			"CNY":          reflect.ValueOf(&q.CNY),
			"Cash":         reflect.ValueOf(&q.Cash),
			"DKK":          reflect.ValueOf(&q.DKK),
			"EUR":          reflect.ValueOf(&q.EUR),
			"GBP":          reflect.ValueOf(&q.GBP),
			"HKD":          reflect.ValueOf(&q.HKD),
			"Historical":   reflect.ValueOf(&q.Historical),
			"IDR":          reflect.ValueOf(&q.IDR),
			"INR":          reflect.ValueOf(&q.INR),
			"ISO":          reflect.ValueOf(&q.ISO),
			"JPY":          reflect.ValueOf(&q.JPY),
			"KRW":          reflect.ValueOf(&q.KRW),
			"MXN":          reflect.ValueOf(&q.MXN),
			"NOK":          reflect.ValueOf(&q.NOK),
			"NZD":          reflect.ValueOf(&q.NZD),
			"NarrowSymbol": reflect.ValueOf(&q.NarrowSymbol),
			"NonTender":    reflect.ValueOf(&q.NonTender),
			"PLN":          reflect.ValueOf(&q.PLN),
			"RUB":          reflect.ValueOf(&q.RUB),
			"SAR":          reflect.ValueOf(&q.SAR),
			"SEK":          reflect.ValueOf(&q.SEK),
			"Standard":     reflect.ValueOf(&q.Standard),
			"Symbol":       reflect.ValueOf(&q.Symbol),
			"THB":          reflect.ValueOf(&q.THB),
			"TRY":          reflect.ValueOf(&q.TRY),
			"TWD":          reflect.ValueOf(&q.TWD),
			"USD":          reflect.ValueOf(&q.USD),
			"XAG":          reflect.ValueOf(&q.XAG),
			"XAU":          reflect.ValueOf(&q.XAU),
			"XPD":          reflect.ValueOf(&q.XPD),
			"XPT":          reflect.ValueOf(&q.XPT),
			"XTS":          reflect.ValueOf(&q.XTS),
			"XXX":          reflect.ValueOf(&q.XXX),
			"ZAR":          reflect.ValueOf(&q.ZAR),
		},
		Funcs: map[string]reflect.Value{
			"Date":         reflect.ValueOf(q.Date),
			"FromRegion":   reflect.ValueOf(q.FromRegion),
			"FromTag":      reflect.ValueOf(q.FromTag),
			"MustParseISO": reflect.ValueOf(q.MustParseISO),
			"ParseISO":     reflect.ValueOf(q.ParseISO),
			"Query":        reflect.ValueOf(q.Query),
			"Region":       reflect.ValueOf(q.Region),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"CLDRVersion": {"untyped string", constant.MakeString(string(q.CLDRVersion))},
		},
	})
}
