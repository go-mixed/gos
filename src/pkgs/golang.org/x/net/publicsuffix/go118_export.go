// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package publicsuffix

import (
	q "golang.org/x/net/publicsuffix"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "publicsuffix",
		Path: "golang.org/x/net/publicsuffix",
		Deps: map[string]string{
			"fmt":                "fmt",
			"net/http/cookiejar": "cookiejar",
			"strings":            "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"List": reflect.ValueOf(&q.List),
		},
		Funcs: map[string]reflect.Value{
			"EffectiveTLDPlusOne": reflect.ValueOf(q.EffectiveTLDPlusOne),
			"PublicSuffix":        reflect.ValueOf(q.PublicSuffix),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
