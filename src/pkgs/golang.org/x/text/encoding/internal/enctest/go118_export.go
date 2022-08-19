// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package enctest

import (
	q "golang.org/x/text/encoding/internal/enctest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "enctest",
		Path: "golang.org/x/text/encoding/internal/enctest",
		Deps: map[string]string{
			"bytes":                      "bytes",
			"fmt":                        "fmt",
			"golang.org/x/text/encoding": "encoding",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
			"io":                                             "io",
			"io/ioutil":                                      "ioutil",
			"strings":                                        "strings",
			"testing":                                        "testing",
		},
		Interfaces: map[string]reflect.Type{
			"Transcoder": reflect.TypeOf((*q.Transcoder)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Benchmark":    reflect.ValueOf(q.Benchmark),
			"TestEncoding": reflect.ValueOf(q.TestEncoding),
			"TestFile":     reflect.ValueOf(q.TestFile),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
