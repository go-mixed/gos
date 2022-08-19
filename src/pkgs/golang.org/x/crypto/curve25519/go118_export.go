// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package curve25519

import (
	q "golang.org/x/crypto/curve25519"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "curve25519",
		Path: "golang.org/x/crypto/curve25519",
		Deps: map[string]string{
			"crypto/subtle": "subtle",
			"errors":        "errors",
			"golang.org/x/crypto/curve25519/internal/field": "field",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Basepoint": reflect.ValueOf(&q.Basepoint),
		},
		Funcs: map[string]reflect.Value{
			"ScalarBaseMult": reflect.ValueOf(q.ScalarBaseMult),
			"ScalarMult":     reflect.ValueOf(q.ScalarMult),
			"X25519":         reflect.ValueOf(q.X25519),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"PointSize":  {"untyped int", constant.MakeInt64(int64(q.PointSize))},
			"ScalarSize": {"untyped int", constant.MakeInt64(int64(q.ScalarSize))},
		},
	})
}
