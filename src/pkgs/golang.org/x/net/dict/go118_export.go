// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package dict

import (
	q "golang.org/x/net/dict"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "dict",
		Path: "golang.org/x/net/dict",
		Deps: map[string]string{
			"net/textproto": "textproto",
			"strconv":       "strconv",
			"strings":       "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Client": reflect.TypeOf((*q.Client)(nil)).Elem(),
			"Defn":   reflect.TypeOf((*q.Defn)(nil)).Elem(),
			"Dict":   reflect.TypeOf((*q.Dict)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dial": reflect.ValueOf(q.Dial),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
