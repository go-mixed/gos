// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package errors

import (
	q "golang.org/x/crypto/openpgp/errors"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "errors",
		Path: "golang.org/x/crypto/openpgp/errors",
		Deps: map[string]string{
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"InvalidArgumentError":   reflect.TypeOf((*q.InvalidArgumentError)(nil)).Elem(),
			"SignatureError":         reflect.TypeOf((*q.SignatureError)(nil)).Elem(),
			"StructuralError":        reflect.TypeOf((*q.StructuralError)(nil)).Elem(),
			"UnknownPacketTypeError": reflect.TypeOf((*q.UnknownPacketTypeError)(nil)).Elem(),
			"UnsupportedError":       reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrKeyIncorrect":  reflect.ValueOf(&q.ErrKeyIncorrect),
			"ErrKeyRevoked":    reflect.ValueOf(&q.ErrKeyRevoked),
			"ErrUnknownIssuer": reflect.ValueOf(&q.ErrUnknownIssuer),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
