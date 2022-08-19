// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bcrypt

import (
	q "golang.org/x/crypto/bcrypt"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bcrypt",
		Path: "golang.org/x/crypto/bcrypt",
		Deps: map[string]string{
			"crypto/rand":                  "rand",
			"crypto/subtle":                "subtle",
			"encoding/base64":              "base64",
			"errors":                       "errors",
			"fmt":                          "fmt",
			"golang.org/x/crypto/blowfish": "blowfish",
			"io":                           "io",
			"strconv":                      "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"HashVersionTooNewError": reflect.TypeOf((*q.HashVersionTooNewError)(nil)).Elem(),
			"InvalidCostError":       reflect.TypeOf((*q.InvalidCostError)(nil)).Elem(),
			"InvalidHashPrefixError": reflect.TypeOf((*q.InvalidHashPrefixError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrHashTooShort":              reflect.ValueOf(&q.ErrHashTooShort),
			"ErrMismatchedHashAndPassword": reflect.ValueOf(&q.ErrMismatchedHashAndPassword),
		},
		Funcs: map[string]reflect.Value{
			"CompareHashAndPassword": reflect.ValueOf(q.CompareHashAndPassword),
			"Cost":                   reflect.ValueOf(q.Cost),
			"GenerateFromPassword":   reflect.ValueOf(q.GenerateFromPassword),
		},
		TypedConsts: map[string]igop.TypedConst{
			"DefaultCost": {reflect.TypeOf(q.DefaultCost), constant.MakeInt64(int64(q.DefaultCost))},
			"MaxCost":     {reflect.TypeOf(q.MaxCost), constant.MakeInt64(int64(q.MaxCost))},
			"MinCost":     {reflect.TypeOf(q.MinCost), constant.MakeInt64(int64(q.MinCost))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
