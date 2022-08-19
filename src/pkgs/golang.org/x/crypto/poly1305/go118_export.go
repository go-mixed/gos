// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package poly1305

import (
	q "golang.org/x/crypto/poly1305"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "poly1305",
		Path: "golang.org/x/crypto/poly1305",
		Deps: map[string]string{
			"golang.org/x/crypto/internal/poly1305": "poly1305",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"MAC": reflect.TypeOf((*q.MAC)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":    reflect.ValueOf(q.New),
			"Sum":    reflect.ValueOf(q.Sum),
			"Verify": reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"TagSize": {"untyped int", constant.MakeInt64(int64(q.TagSize))},
		},
	})
}
