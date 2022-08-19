// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package rate

import (
	q "golang.org/x/time/rate"

	"go/constant"
	"go/token"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rate",
		Path: "golang.org/x/time/rate",
		Deps: map[string]string{
			"context": "context",
			"fmt":     "fmt",
			"math":    "math",
			"sync":    "sync",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Limit":       reflect.TypeOf((*q.Limit)(nil)).Elem(),
			"Limiter":     reflect.TypeOf((*q.Limiter)(nil)).Elem(),
			"Reservation": reflect.TypeOf((*q.Reservation)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Every":      reflect.ValueOf(q.Every),
			"NewLimiter": reflect.ValueOf(q.NewLimiter),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Inf":         {reflect.TypeOf(q.Inf), constant.MakeFromLiteral("1.79769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368e+308", token.FLOAT, 0)},
			"InfDuration": {reflect.TypeOf(q.InfDuration), constant.MakeInt64(int64(q.InfDuration))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
