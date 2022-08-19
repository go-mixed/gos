// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package websocket

import (
	q "golang.org/x/net/websocket"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "websocket",
		Path: "golang.org/x/net/websocket",
		Deps: map[string]string{
			"bufio":           "bufio",
			"bytes":           "bytes",
			"crypto/rand":     "rand",
			"crypto/sha1":     "sha1",
			"crypto/tls":      "tls",
			"encoding/base64": "base64",
			"encoding/binary": "binary",
			"encoding/json":   "json",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"io/ioutil":       "ioutil",
			"net":             "net",
			"net/http":        "http",
			"net/url":         "url",
			"strings":         "strings",
			"sync":            "sync",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Addr":          reflect.TypeOf((*q.Addr)(nil)).Elem(),
			"Codec":         reflect.TypeOf((*q.Codec)(nil)).Elem(),
			"Config":        reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Conn":          reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"DialError":     reflect.TypeOf((*q.DialError)(nil)).Elem(),
			"Handler":       reflect.TypeOf((*q.Handler)(nil)).Elem(),
			"ProtocolError": reflect.TypeOf((*q.ProtocolError)(nil)).Elem(),
			"Server":        reflect.TypeOf((*q.Server)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrBadClosingStatus":      reflect.ValueOf(&q.ErrBadClosingStatus),
			"ErrBadFrame":              reflect.ValueOf(&q.ErrBadFrame),
			"ErrBadFrameBoundary":      reflect.ValueOf(&q.ErrBadFrameBoundary),
			"ErrBadMaskingKey":         reflect.ValueOf(&q.ErrBadMaskingKey),
			"ErrBadPongMessage":        reflect.ValueOf(&q.ErrBadPongMessage),
			"ErrBadProtocolVersion":    reflect.ValueOf(&q.ErrBadProtocolVersion),
			"ErrBadRequestMethod":      reflect.ValueOf(&q.ErrBadRequestMethod),
			"ErrBadScheme":             reflect.ValueOf(&q.ErrBadScheme),
			"ErrBadStatus":             reflect.ValueOf(&q.ErrBadStatus),
			"ErrBadUpgrade":            reflect.ValueOf(&q.ErrBadUpgrade),
			"ErrBadWebSocketLocation":  reflect.ValueOf(&q.ErrBadWebSocketLocation),
			"ErrBadWebSocketOrigin":    reflect.ValueOf(&q.ErrBadWebSocketOrigin),
			"ErrBadWebSocketProtocol":  reflect.ValueOf(&q.ErrBadWebSocketProtocol),
			"ErrBadWebSocketVersion":   reflect.ValueOf(&q.ErrBadWebSocketVersion),
			"ErrChallengeResponse":     reflect.ValueOf(&q.ErrChallengeResponse),
			"ErrFrameTooLarge":         reflect.ValueOf(&q.ErrFrameTooLarge),
			"ErrNotImplemented":        reflect.ValueOf(&q.ErrNotImplemented),
			"ErrNotSupported":          reflect.ValueOf(&q.ErrNotSupported),
			"ErrNotWebSocket":          reflect.ValueOf(&q.ErrNotWebSocket),
			"ErrUnsupportedExtensions": reflect.ValueOf(&q.ErrUnsupportedExtensions),
			"JSON":                     reflect.ValueOf(&q.JSON),
			"Message":                  reflect.ValueOf(&q.Message),
		},
		Funcs: map[string]reflect.Value{
			"Dial":       reflect.ValueOf(q.Dial),
			"DialConfig": reflect.ValueOf(q.DialConfig),
			"NewClient":  reflect.ValueOf(q.NewClient),
			"NewConfig":  reflect.ValueOf(q.NewConfig),
			"Origin":     reflect.ValueOf(q.Origin),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BinaryFrame":              {"untyped int", constant.MakeInt64(int64(q.BinaryFrame))},
			"CloseFrame":               {"untyped int", constant.MakeInt64(int64(q.CloseFrame))},
			"ContinuationFrame":        {"untyped int", constant.MakeInt64(int64(q.ContinuationFrame))},
			"DefaultMaxPayloadBytes":   {"untyped int", constant.MakeInt64(int64(q.DefaultMaxPayloadBytes))},
			"PingFrame":                {"untyped int", constant.MakeInt64(int64(q.PingFrame))},
			"PongFrame":                {"untyped int", constant.MakeInt64(int64(q.PongFrame))},
			"ProtocolVersionHybi":      {"untyped int", constant.MakeInt64(int64(q.ProtocolVersionHybi))},
			"ProtocolVersionHybi13":    {"untyped int", constant.MakeInt64(int64(q.ProtocolVersionHybi13))},
			"SupportedProtocolVersion": {"untyped string", constant.MakeString(string(q.SupportedProtocolVersion))},
			"TextFrame":                {"untyped int", constant.MakeInt64(int64(q.TextFrame))},
			"UnknownFrame":             {"untyped int", constant.MakeInt64(int64(q.UnknownFrame))},
		},
	})
}
