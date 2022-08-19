// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package blake2s

import (
	q "golang.org/x/crypto/blake2s"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "blake2s",
		Path: "golang.org/x/crypto/blake2s",
		Deps: map[string]string{
			"crypto":               "crypto",
			"encoding/binary":      "binary",
			"errors":               "errors",
			"golang.org/x/sys/cpu": "cpu",
			"hash":                 "hash",
			"io":                   "io",
			"math/bits":            "bits",
		},
		Interfaces: map[string]reflect.Type{
			"XOF": reflect.TypeOf((*q.XOF)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New128": reflect.ValueOf(q.New128),
			"New256": reflect.ValueOf(q.New256),
			"NewXOF": reflect.ValueOf(q.NewXOF),
			"Sum256": reflect.ValueOf(q.Sum256),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize":           {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"OutputLengthUnknown": {"untyped int", constant.MakeInt64(int64(q.OutputLengthUnknown))},
			"Size":                {"untyped int", constant.MakeInt64(int64(q.Size))},
			"Size128":             {"untyped int", constant.MakeInt64(int64(q.Size128))},
		},
	})
}
