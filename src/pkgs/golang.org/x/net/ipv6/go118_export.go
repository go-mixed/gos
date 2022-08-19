// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ipv6

import (
	q "golang.org/x/net/ipv6"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ipv6",
		Path: "golang.org/x/net/ipv6",
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
			"ICMPFilter":     reflect.TypeOf((*q.ICMPFilter)(nil)).Elem(),
			"ICMPType":       reflect.TypeOf((*q.ICMPType)(nil)).Elem(),
			"PacketConn":     reflect.TypeOf((*q.PacketConn)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Message": reflect.TypeOf((*q.Message)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewConn":           reflect.ValueOf(q.NewConn),
			"NewControlMessage": reflect.ValueOf(q.NewControlMessage),
			"NewPacketConn":     reflect.ValueOf(q.NewPacketConn),
			"ParseHeader":       reflect.ValueOf(q.ParseHeader),
		},
		TypedConsts: map[string]igop.TypedConst{
			"FlagDst":                                       {reflect.TypeOf(q.FlagDst), constant.MakeInt64(int64(q.FlagDst))},
			"FlagHopLimit":                                  {reflect.TypeOf(q.FlagHopLimit), constant.MakeInt64(int64(q.FlagHopLimit))},
			"FlagInterface":                                 {reflect.TypeOf(q.FlagInterface), constant.MakeInt64(int64(q.FlagInterface))},
			"FlagPathMTU":                                   {reflect.TypeOf(q.FlagPathMTU), constant.MakeInt64(int64(q.FlagPathMTU))},
			"FlagSrc":                                       {reflect.TypeOf(q.FlagSrc), constant.MakeInt64(int64(q.FlagSrc))},
			"FlagTrafficClass":                              {reflect.TypeOf(q.FlagTrafficClass), constant.MakeInt64(int64(q.FlagTrafficClass))},
			"ICMPTypeCertificationPathAdvertisement":        {reflect.TypeOf(q.ICMPTypeCertificationPathAdvertisement), constant.MakeInt64(int64(q.ICMPTypeCertificationPathAdvertisement))},
			"ICMPTypeCertificationPathSolicitation":         {reflect.TypeOf(q.ICMPTypeCertificationPathSolicitation), constant.MakeInt64(int64(q.ICMPTypeCertificationPathSolicitation))},
			"ICMPTypeDestinationUnreachable":                {reflect.TypeOf(q.ICMPTypeDestinationUnreachable), constant.MakeInt64(int64(q.ICMPTypeDestinationUnreachable))},
			"ICMPTypeDuplicateAddressConfirmation":          {reflect.TypeOf(q.ICMPTypeDuplicateAddressConfirmation), constant.MakeInt64(int64(q.ICMPTypeDuplicateAddressConfirmation))},
			"ICMPTypeDuplicateAddressRequest":               {reflect.TypeOf(q.ICMPTypeDuplicateAddressRequest), constant.MakeInt64(int64(q.ICMPTypeDuplicateAddressRequest))},
			"ICMPTypeEchoReply":                             {reflect.TypeOf(q.ICMPTypeEchoReply), constant.MakeInt64(int64(q.ICMPTypeEchoReply))},
			"ICMPTypeEchoRequest":                           {reflect.TypeOf(q.ICMPTypeEchoRequest), constant.MakeInt64(int64(q.ICMPTypeEchoRequest))},
			"ICMPTypeExtendedEchoReply":                     {reflect.TypeOf(q.ICMPTypeExtendedEchoReply), constant.MakeInt64(int64(q.ICMPTypeExtendedEchoReply))},
			"ICMPTypeExtendedEchoRequest":                   {reflect.TypeOf(q.ICMPTypeExtendedEchoRequest), constant.MakeInt64(int64(q.ICMPTypeExtendedEchoRequest))},
			"ICMPTypeFMIPv6":                                {reflect.TypeOf(q.ICMPTypeFMIPv6), constant.MakeInt64(int64(q.ICMPTypeFMIPv6))},
			"ICMPTypeHomeAgentAddressDiscoveryReply":        {reflect.TypeOf(q.ICMPTypeHomeAgentAddressDiscoveryReply), constant.MakeInt64(int64(q.ICMPTypeHomeAgentAddressDiscoveryReply))},
			"ICMPTypeHomeAgentAddressDiscoveryRequest":      {reflect.TypeOf(q.ICMPTypeHomeAgentAddressDiscoveryRequest), constant.MakeInt64(int64(q.ICMPTypeHomeAgentAddressDiscoveryRequest))},
			"ICMPTypeILNPv6LocatorUpdate":                   {reflect.TypeOf(q.ICMPTypeILNPv6LocatorUpdate), constant.MakeInt64(int64(q.ICMPTypeILNPv6LocatorUpdate))},
			"ICMPTypeInverseNeighborDiscoveryAdvertisement": {reflect.TypeOf(q.ICMPTypeInverseNeighborDiscoveryAdvertisement), constant.MakeInt64(int64(q.ICMPTypeInverseNeighborDiscoveryAdvertisement))},
			"ICMPTypeInverseNeighborDiscoverySolicitation":  {reflect.TypeOf(q.ICMPTypeInverseNeighborDiscoverySolicitation), constant.MakeInt64(int64(q.ICMPTypeInverseNeighborDiscoverySolicitation))},
			"ICMPTypeMPLControl":                            {reflect.TypeOf(q.ICMPTypeMPLControl), constant.MakeInt64(int64(q.ICMPTypeMPLControl))},
			"ICMPTypeMobilePrefixAdvertisement":             {reflect.TypeOf(q.ICMPTypeMobilePrefixAdvertisement), constant.MakeInt64(int64(q.ICMPTypeMobilePrefixAdvertisement))},
			"ICMPTypeMobilePrefixSolicitation":              {reflect.TypeOf(q.ICMPTypeMobilePrefixSolicitation), constant.MakeInt64(int64(q.ICMPTypeMobilePrefixSolicitation))},
			"ICMPTypeMulticastListenerDone":                 {reflect.TypeOf(q.ICMPTypeMulticastListenerDone), constant.MakeInt64(int64(q.ICMPTypeMulticastListenerDone))},
			"ICMPTypeMulticastListenerQuery":                {reflect.TypeOf(q.ICMPTypeMulticastListenerQuery), constant.MakeInt64(int64(q.ICMPTypeMulticastListenerQuery))},
			"ICMPTypeMulticastListenerReport":               {reflect.TypeOf(q.ICMPTypeMulticastListenerReport), constant.MakeInt64(int64(q.ICMPTypeMulticastListenerReport))},
			"ICMPTypeMulticastRouterAdvertisement":          {reflect.TypeOf(q.ICMPTypeMulticastRouterAdvertisement), constant.MakeInt64(int64(q.ICMPTypeMulticastRouterAdvertisement))},
			"ICMPTypeMulticastRouterSolicitation":           {reflect.TypeOf(q.ICMPTypeMulticastRouterSolicitation), constant.MakeInt64(int64(q.ICMPTypeMulticastRouterSolicitation))},
			"ICMPTypeMulticastRouterTermination":            {reflect.TypeOf(q.ICMPTypeMulticastRouterTermination), constant.MakeInt64(int64(q.ICMPTypeMulticastRouterTermination))},
			"ICMPTypeNeighborAdvertisement":                 {reflect.TypeOf(q.ICMPTypeNeighborAdvertisement), constant.MakeInt64(int64(q.ICMPTypeNeighborAdvertisement))},
			"ICMPTypeNeighborSolicitation":                  {reflect.TypeOf(q.ICMPTypeNeighborSolicitation), constant.MakeInt64(int64(q.ICMPTypeNeighborSolicitation))},
			"ICMPTypeNodeInformationQuery":                  {reflect.TypeOf(q.ICMPTypeNodeInformationQuery), constant.MakeInt64(int64(q.ICMPTypeNodeInformationQuery))},
			"ICMPTypeNodeInformationResponse":               {reflect.TypeOf(q.ICMPTypeNodeInformationResponse), constant.MakeInt64(int64(q.ICMPTypeNodeInformationResponse))},
			"ICMPTypePacketTooBig":                          {reflect.TypeOf(q.ICMPTypePacketTooBig), constant.MakeInt64(int64(q.ICMPTypePacketTooBig))},
			"ICMPTypeParameterProblem":                      {reflect.TypeOf(q.ICMPTypeParameterProblem), constant.MakeInt64(int64(q.ICMPTypeParameterProblem))},
			"ICMPTypeRPLControl":                            {reflect.TypeOf(q.ICMPTypeRPLControl), constant.MakeInt64(int64(q.ICMPTypeRPLControl))},
			"ICMPTypeRedirect":                              {reflect.TypeOf(q.ICMPTypeRedirect), constant.MakeInt64(int64(q.ICMPTypeRedirect))},
			"ICMPTypeRouterAdvertisement":                   {reflect.TypeOf(q.ICMPTypeRouterAdvertisement), constant.MakeInt64(int64(q.ICMPTypeRouterAdvertisement))},
			"ICMPTypeRouterRenumbering":                     {reflect.TypeOf(q.ICMPTypeRouterRenumbering), constant.MakeInt64(int64(q.ICMPTypeRouterRenumbering))},
			"ICMPTypeRouterSolicitation":                    {reflect.TypeOf(q.ICMPTypeRouterSolicitation), constant.MakeInt64(int64(q.ICMPTypeRouterSolicitation))},
			"ICMPTypeTimeExceeded":                          {reflect.TypeOf(q.ICMPTypeTimeExceeded), constant.MakeInt64(int64(q.ICMPTypeTimeExceeded))},
			"ICMPTypeVersion2MulticastListenerReport":       {reflect.TypeOf(q.ICMPTypeVersion2MulticastListenerReport), constant.MakeInt64(int64(q.ICMPTypeVersion2MulticastListenerReport))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"HeaderLen": {"untyped int", constant.MakeInt64(int64(q.HeaderLen))},
			"Version":   {"untyped int", constant.MakeInt64(int64(q.Version))},
		},
	})
}
