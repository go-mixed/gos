// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package proxy

import (
	q "golang.org/x/net/proxy"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "proxy",
		Path: "golang.org/x/net/proxy",
		Deps: map[string]string{
			"context":                         "context",
			"errors":                          "errors",
			"golang.org/x/net/internal/socks": "socks",
			"net":                             "net",
			"net/url":                         "url",
			"os":                              "os",
			"strings":                         "strings",
			"sync":                            "sync",
		},
		Interfaces: map[string]reflect.Type{
			"ContextDialer": reflect.TypeOf((*q.ContextDialer)(nil)).Elem(),
			"Dialer":        reflect.TypeOf((*q.Dialer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Auth":    reflect.TypeOf((*q.Auth)(nil)).Elem(),
			"PerHost": reflect.TypeOf((*q.PerHost)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Direct": reflect.ValueOf(&q.Direct),
		},
		Funcs: map[string]reflect.Value{
			"Dial":                 reflect.ValueOf(q.Dial),
			"FromEnvironment":      reflect.ValueOf(q.FromEnvironment),
			"FromEnvironmentUsing": reflect.ValueOf(q.FromEnvironmentUsing),
			"FromURL":              reflect.ValueOf(q.FromURL),
			"NewPerHost":           reflect.ValueOf(q.NewPerHost),
			"RegisterDialerType":   reflect.ValueOf(q.RegisterDialerType),
			"SOCKS5":               reflect.ValueOf(q.SOCKS5),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
