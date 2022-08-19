// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ipv4

import (
	q "golang.org/x/net/ipv4"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ipv4",
		Path: "golang.org/x/net/ipv4",
		Deps: map[string]string{
			"encoding/binary":                  "binary",
			"errors":                           "errors",
			"fmt":                              "fmt",
			"golang.org/x/net/bpf":             "bpf",
			"golang.org/x/net/internal/iana":   "iana",
			"golang.org/x/net/internal/socket": "socket",
			"golang.org/x/sys/unix":            "unix",
			"net":                              "net",
			"runtime":                          "runtime",
			"sync":                             "sync",
			"syscall":                          "syscall",
			"time":                             "time",
			"unsafe":                           "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Conn":           reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"ControlFlags":   reflect.TypeOf((*q.ControlFlags)(nil)).Elem(),
			"ControlMessage": reflect.TypeOf((*q.ControlMessage)(nil)).Elem(),
			"Header":         reflect.TypeOf((*q.Header)(nil)).Elem(),
			"HeaderFlags":    reflect.TypeOf((*q.HeaderFlags)(nil)).Elem(),
			"ICMPFilter":     reflect.TypeOf((*q.ICMPFilter)(nil)).Elem(),
			"ICMPType":       reflect.TypeOf((*q.ICMPType)(nil)).Elem(),
			"PacketConn":     reflect.TypeOf((*q.PacketConn)(nil)).Elem(),
			"RawConn":        reflect.TypeOf((*q.RawConn)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Message": reflect.TypeOf((*q.Message)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewConn":           reflect.ValueOf(q.NewConn),
			"NewControlMessage": reflect.ValueOf(q.NewControlMessage),
			"NewPacketConn":     reflect.ValueOf(q.NewPacketConn),
			"NewRawConn":        reflect.ValueOf(q.NewRawConn),
			"ParseHeader":       reflect.ValueOf(q.ParseHeader),
		},
		TypedConsts: map[string]igop.TypedConst{
			"DontFragment":                   {reflect.TypeOf(q.DontFragment), constant.MakeInt64(int64(q.DontFragment))},
			"FlagDst":                        {reflect.TypeOf(q.FlagDst), constant.MakeInt64(int64(q.FlagDst))},
			"FlagInterface":                  {reflect.TypeOf(q.FlagInterface), constant.MakeInt64(int64(q.FlagInterface))},
			"FlagSrc":                        {reflect.TypeOf(q.FlagSrc), constant.MakeInt64(int64(q.FlagSrc))},
			"FlagTTL":                        {reflect.TypeOf(q.FlagTTL), constant.MakeInt64(int64(q.FlagTTL))},
			"ICMPTypeDestinationUnreachable": {reflect.TypeOf(q.ICMPTypeDestinationUnreachable), constant.MakeInt64(int64(q.ICMPTypeDestinationUnreachable))},
			"ICMPTypeEcho":                   {reflect.TypeOf(q.ICMPTypeEcho), constant.MakeInt64(int64(q.ICMPTypeEcho))},
			"ICMPTypeEchoReply":              {reflect.TypeOf(q.ICMPTypeEchoReply), constant.MakeInt64(int64(q.ICMPTypeEchoReply))},
			"ICMPTypeExtendedEchoReply":      {reflect.TypeOf(q.ICMPTypeExtendedEchoReply), constant.MakeInt64(int64(q.ICMPTypeExtendedEchoReply))},
			"ICMPTypeExtendedEchoRequest":    {reflect.TypeOf(q.ICMPTypeExtendedEchoRequest), constant.MakeInt64(int64(q.ICMPTypeExtendedEchoRequest))},
			"ICMPTypeParameterProblem":       {reflect.TypeOf(q.ICMPTypeParameterProblem), constant.MakeInt64(int64(q.ICMPTypeParameterProblem))},
			"ICMPTypePhoturis":               {reflect.TypeOf(q.ICMPTypePhoturis), constant.MakeInt64(int64(q.ICMPTypePhoturis))},
			"ICMPTypeRedirect":               {reflect.TypeOf(q.ICMPTypeRedirect), constant.MakeInt64(int64(q.ICMPTypeRedirect))},
			"ICMPTypeRouterAdvertisement":    {reflect.TypeOf(q.ICMPTypeRouterAdvertisement), constant.MakeInt64(int64(q.ICMPTypeRouterAdvertisement))},
			"ICMPTypeRouterSolicitation":     {reflect.TypeOf(q.ICMPTypeRouterSolicitation), constant.MakeInt64(int64(q.ICMPTypeRouterSolicitation))},
			"ICMPTypeTimeExceeded":           {reflect.TypeOf(q.ICMPTypeTimeExceeded), constant.MakeInt64(int64(q.ICMPTypeTimeExceeded))},
			"ICMPTypeTimestamp":              {reflect.TypeOf(q.ICMPTypeTimestamp), constant.MakeInt64(int64(q.ICMPTypeTimestamp))},
			"ICMPTypeTimestampReply":         {reflect.TypeOf(q.ICMPTypeTimestampReply), constant.MakeInt64(int64(q.ICMPTypeTimestampReply))},
			"MoreFragments":                  {reflect.TypeOf(q.MoreFragments), constant.MakeInt64(int64(q.MoreFragments))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"HeaderLen": {"untyped int", constant.MakeInt64(int64(q.HeaderLen))},
			"Version":   {"untyped int", constant.MakeInt64(int64(q.Version))},
		},
	})
}
