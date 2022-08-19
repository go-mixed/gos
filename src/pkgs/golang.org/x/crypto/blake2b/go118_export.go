// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package blake2b

import (
	q "golang.org/x/crypto/blake2b"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "blake2b",
		Path: "golang.org/x/crypto/blake2b",
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
			"New":    reflect.ValueOf(q.New),
			"New256": reflect.ValueOf(q.New256),
			"New384": reflect.ValueOf(q.New384),
			"New512": reflect.ValueOf(q.New512),
			"NewXOF": reflect.ValueOf(q.NewXOF),
			"Sum256": reflect.ValueOf(q.Sum256),
			"Sum384": reflect.ValueOf(q.Sum384),
			"Sum512": reflect.ValueOf(q.Sum512),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize":           {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"OutputLengthUnknown": {"untyped int", constant.MakeInt64(int64(q.OutputLengthUnknown))},
			"Size":                {"untyped int", constant.MakeInt64(int64(q.Size))},
			"Size256":             {"untyped int", constant.MakeInt64(int64(q.Size256))},
			"Size384":             {"untyped int", constant.MakeInt64(int64(q.Size384))},
		},
	})
}
