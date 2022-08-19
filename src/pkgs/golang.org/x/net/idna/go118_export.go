// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package idna

import (
	q "golang.org/x/net/idna"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "idna",
		Path: "golang.org/x/net/idna",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"golang.org/x/text/secure/bidirule": "bidirule",
			"golang.org/x/text/unicode/bidi":    "bidi",
			"golang.org/x/text/unicode/norm":    "norm",
			"math":                              "math",
			"strings":                           "strings",
			"unicode/utf8":                      "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Option":  reflect.TypeOf((*q.Option)(nil)).Elem(),
			"Profile": reflect.TypeOf((*q.Profile)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Display":      reflect.ValueOf(&q.Display),
			"Lookup":       reflect.ValueOf(&q.Lookup),
			"Punycode":     reflect.ValueOf(&q.Punycode),
			"Registration": reflect.ValueOf(&q.Registration),
		},
		Funcs: map[string]reflect.Value{
			"BidiRule":                reflect.ValueOf(q.BidiRule),
			"CheckHyphens":            reflect.ValueOf(q.CheckHyphens),
			"CheckJoiners":            reflect.ValueOf(q.CheckJoiners),
			"MapForLookup":            reflect.ValueOf(q.MapForLookup),
			"New":                     reflect.ValueOf(q.New),
			"RemoveLeadingDots":       reflect.ValueOf(q.RemoveLeadingDots),
			"StrictDomainName":        reflect.ValueOf(q.StrictDomainName),
			"ToASCII":                 reflect.ValueOf(q.ToASCII),
			"ToUnicode":               reflect.ValueOf(q.ToUnicode),
			"Transitional":            reflect.ValueOf(q.Transitional),
			"ValidateForRegistration": reflect.ValueOf(q.ValidateForRegistration),
			"ValidateLabels":          reflect.ValueOf(q.ValidateLabels),
			"VerifyDNSLength":         reflect.ValueOf(q.VerifyDNSLength),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
