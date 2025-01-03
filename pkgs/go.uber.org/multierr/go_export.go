// export by github.com/goplus/igop/cmd/qexp

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
			"bytes":       "bytes",
			"errors":      "errors",
			"fmt":         "fmt",
			"io":          "io",
			"strings":     "strings",
			"sync":        "sync",
			"sync/atomic": "atomic",
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
			"AppendFunc":   reflect.ValueOf(q.AppendFunc),
			"AppendInto":   reflect.ValueOf(q.AppendInto),
			"AppendInvoke": reflect.ValueOf(q.AppendInvoke),
			"Close":        reflect.ValueOf(q.Close),
			"Combine":      reflect.ValueOf(q.Combine),
			"Errors":       reflect.ValueOf(q.Errors),
			"Every":        reflect.ValueOf(q.Every),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
