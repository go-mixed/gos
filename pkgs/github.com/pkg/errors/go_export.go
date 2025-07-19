// export by github.com/goplus/ixgo/cmd/qexp

package errors

import (
	"github.com/goplus/ixgo"
	q "github.com/pkg/errors"

	"reflect"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "errors",
		Path: "github.com/pkg/errors",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"path":    "path",
			"runtime": "runtime",
			"strconv": "strconv",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Frame":      reflect.TypeOf((*q.Frame)(nil)).Elem(),
			"StackTrace": reflect.TypeOf((*q.StackTrace)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"As":           reflect.ValueOf(q.As),
			"Cause":        reflect.ValueOf(q.Cause),
			"Errorf":       reflect.ValueOf(q.Errorf),
			"Is":           reflect.ValueOf(q.Is),
			"New":          reflect.ValueOf(q.New),
			"Unwrap":       reflect.ValueOf(q.Unwrap),
			"WithMessage":  reflect.ValueOf(q.WithMessage),
			"WithMessagef": reflect.ValueOf(q.WithMessagef),
			"WithStack":    reflect.ValueOf(q.WithStack),
			"Wrap":         reflect.ValueOf(q.Wrap),
			"Wrapf":        reflect.ValueOf(q.Wrapf),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
