// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package trace

import (
	q "golang.org/x/net/trace"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "trace",
		Path: "golang.org/x/net/trace",
		Deps: map[string]string{
			"bytes":                                "bytes",
			"context":                              "context",
			"fmt":                                  "fmt",
			"golang.org/x/net/internal/timeseries": "timeseries",
			"html/template":                        "template",
			"io":                                   "io",
			"log":                                  "log",
			"math":                                 "math",
			"net":                                  "net",
			"net/http":                             "http",
			"net/url":                              "url",
			"runtime":                              "runtime",
			"sort":                                 "sort",
			"strconv":                              "strconv",
			"strings":                              "strings",
			"sync":                                 "sync",
			"sync/atomic":                          "atomic",
			"text/tabwriter":                       "tabwriter",
			"time":                                 "time",
		},
		Interfaces: map[string]reflect.Type{
			"EventLog": reflect.TypeOf((*q.EventLog)(nil)).Elem(),
			"Trace":    reflect.TypeOf((*q.Trace)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"AuthRequest":         reflect.ValueOf(&q.AuthRequest),
			"DebugUseAfterFinish": reflect.ValueOf(&q.DebugUseAfterFinish),
		},
		Funcs: map[string]reflect.Value{
			"Events":       reflect.ValueOf(q.Events),
			"FromContext":  reflect.ValueOf(q.FromContext),
			"New":          reflect.ValueOf(q.New),
			"NewContext":   reflect.ValueOf(q.NewContext),
			"NewEventLog":  reflect.ValueOf(q.NewEventLog),
			"Render":       reflect.ValueOf(q.Render),
			"RenderEvents": reflect.ValueOf(q.RenderEvents),
			"Traces":       reflect.ValueOf(q.Traces),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
