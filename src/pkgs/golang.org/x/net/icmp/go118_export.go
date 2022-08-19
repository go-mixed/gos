// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package icmp

import (
	q "golang.org/x/net/icmp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "icmp",
		Path: "golang.org/x/net/icmp",
		Deps: map[string]string{
			"encoding/binary":                  "binary",
			"errors":                           "errors",
			"golang.org/x/net/internal/iana":   "iana",
			"golang.org/x/net/internal/socket": "socket",
			"golang.org/x/net/ipv4":            "ipv4",
			"golang.org/x/net/ipv6":            "ipv6",
			"net":                              "net",
			"os":                               "os",
			"runtime":                          "runtime",
			"strconv":                          "strconv",
			"strings":                          "strings",
			"syscall":                          "syscall",
			"time":                             "time",
		},
		Interfaces: map[string]reflect.Type{
			"Extension":   reflect.TypeOf((*q.Extension)(nil)).Elem(),
			"MessageBody": reflect.TypeOf((*q.MessageBody)(nil)).Elem(),
			"Type":        reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"DstUnreach":          reflect.TypeOf((*q.DstUnreach)(nil)).Elem(),
			"Echo":                reflect.TypeOf((*q.Echo)(nil)).Elem(),
			"ExtendedEchoReply":   reflect.TypeOf((*q.ExtendedEchoReply)(nil)).Elem(),
			"ExtendedEchoRequest": reflect.TypeOf((*q.ExtendedEchoRequest)(nil)).Elem(),
			"InterfaceIdent":      reflect.TypeOf((*q.InterfaceIdent)(nil)).Elem(),
			"InterfaceInfo":       reflect.TypeOf((*q.InterfaceInfo)(nil)).Elem(),
			"MPLSLabel":           reflect.TypeOf((*q.MPLSLabel)(nil)).Elem(),
			"MPLSLabelStack":      reflect.TypeOf((*q.MPLSLabelStack)(nil)).Elem(),
			"Message":             reflect.TypeOf((*q.Message)(nil)).Elem(),
			"PacketConn":          reflect.TypeOf((*q.PacketConn)(nil)).Elem(),
			"PacketTooBig":        reflect.TypeOf((*q.PacketTooBig)(nil)).Elem(),
			"ParamProb":           reflect.TypeOf((*q.ParamProb)(nil)).Elem(),
			"RawBody":             reflect.TypeOf((*q.RawBody)(nil)).Elem(),
			"RawExtension":        reflect.TypeOf((*q.RawExtension)(nil)).Elem(),
			"TimeExceeded":        reflect.TypeOf((*q.TimeExceeded)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"DefaultMessageBody": reflect.TypeOf((*q.DefaultMessageBody)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IPv6PseudoHeader": reflect.ValueOf(q.IPv6PseudoHeader),
			"ListenPacket":     reflect.ValueOf(q.ListenPacket),
			"ParseIPv4Header":  reflect.ValueOf(q.ParseIPv4Header),
			"ParseMessage":     reflect.ValueOf(q.ParseMessage),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
