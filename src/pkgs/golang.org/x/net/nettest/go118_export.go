// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package nettest

import (
	q "golang.org/x/net/nettest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "nettest",
		Path: "golang.org/x/net/nettest",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"io/ioutil":       "ioutil",
			"math/rand":       "rand",
			"net":             "net",
			"os":              "os",
			"os/exec":         "exec",
			"runtime":         "runtime",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
			"syscall":         "syscall",
			"testing":         "testing",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"MakePipe": reflect.TypeOf((*q.MakePipe)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"LocalPath":              reflect.ValueOf(q.LocalPath),
			"LoopbackInterface":      reflect.ValueOf(q.LoopbackInterface),
			"MulticastSource":        reflect.ValueOf(q.MulticastSource),
			"NewLocalListener":       reflect.ValueOf(q.NewLocalListener),
			"NewLocalPacketListener": reflect.ValueOf(q.NewLocalPacketListener),
			"RoutedInterface":        reflect.ValueOf(q.RoutedInterface),
			"SupportsIPv4":           reflect.ValueOf(q.SupportsIPv4),
			"SupportsIPv6":           reflect.ValueOf(q.SupportsIPv6),
			"SupportsRawSocket":      reflect.ValueOf(q.SupportsRawSocket),
			"TestConn":               reflect.ValueOf(q.TestConn),
			"TestableAddress":        reflect.ValueOf(q.TestableAddress),
			"TestableNetwork":        reflect.ValueOf(q.TestableNetwork),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
