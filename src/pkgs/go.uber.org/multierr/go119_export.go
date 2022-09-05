// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19
// +build go1.19

package multierr

import (
	q "go.uber.org/multierr"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "multierr",
		Path: "go.uber.org/multierr",
		Deps: map[string]string{
			"bytes":              "bytes",
			"errors":             "errors",
			"fmt":                "fmt",
			"go.uber.org/atomic": "atomic",
			"io":                 "io",
			"strings":            "strings",
			"sync":               "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Invoker": reflect.TypeOf((*q.Invoker)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Invoke": reflect.TypeOf((*q.Invoke)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Append":       reflect.ValueOf(q.Append),
			"AppendInto":   reflect.ValueOf(q.AppendInto),
			"AppendInvoke": reflect.ValueOf(q.AppendInvoke),
			"Close":        reflect.ValueOf(q.Close),
			"Combine":      reflect.ValueOf(q.Combine),
			"Errors":       reflect.ValueOf(q.Errors),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
