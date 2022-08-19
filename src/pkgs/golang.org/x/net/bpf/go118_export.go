// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bpf

import (
	q "golang.org/x/net/bpf"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bpf",
		Path: "golang.org/x/net/bpf",
		Deps: map[string]string{
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
		},
		Interfaces: map[string]reflect.Type{
			"Instruction": reflect.TypeOf((*q.Instruction)(nil)).Elem(),
			"Setter":      reflect.TypeOf((*q.Setter)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ALUOp":          reflect.TypeOf((*q.ALUOp)(nil)).Elem(),
			"ALUOpConstant":  reflect.TypeOf((*q.ALUOpConstant)(nil)).Elem(),
			"ALUOpX":         reflect.TypeOf((*q.ALUOpX)(nil)).Elem(),
			"Extension":      reflect.TypeOf((*q.Extension)(nil)).Elem(),
			"Jump":           reflect.TypeOf((*q.Jump)(nil)).Elem(),
			"JumpIf":         reflect.TypeOf((*q.JumpIf)(nil)).Elem(),
			"JumpIfX":        reflect.TypeOf((*q.JumpIfX)(nil)).Elem(),
			"JumpTest":       reflect.TypeOf((*q.JumpTest)(nil)).Elem(),
			"LoadAbsolute":   reflect.TypeOf((*q.LoadAbsolute)(nil)).Elem(),
			"LoadConstant":   reflect.TypeOf((*q.LoadConstant)(nil)).Elem(),
			"LoadExtension":  reflect.TypeOf((*q.LoadExtension)(nil)).Elem(),
			"LoadIndirect":   reflect.TypeOf((*q.LoadIndirect)(nil)).Elem(),
			"LoadMemShift":   reflect.TypeOf((*q.LoadMemShift)(nil)).Elem(),
			"LoadScratch":    reflect.TypeOf((*q.LoadScratch)(nil)).Elem(),
			"NegateA":        reflect.TypeOf((*q.NegateA)(nil)).Elem(),
			"RawInstruction": reflect.TypeOf((*q.RawInstruction)(nil)).Elem(),
			"Register":       reflect.TypeOf((*q.Register)(nil)).Elem(),
			"RetA":           reflect.TypeOf((*q.RetA)(nil)).Elem(),
			"RetConstant":    reflect.TypeOf((*q.RetConstant)(nil)).Elem(),
			"StoreScratch":   reflect.TypeOf((*q.StoreScratch)(nil)).Elem(),
			"TAX":            reflect.TypeOf((*q.TAX)(nil)).Elem(),
			"TXA":            reflect.TypeOf((*q.TXA)(nil)).Elem(),
			"VM":             reflect.TypeOf((*q.VM)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Assemble":    reflect.ValueOf(q.Assemble),
			"Disassemble": reflect.ValueOf(q.Disassemble),
			"NewVM":       reflect.ValueOf(q.NewVM),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ALUOpAdd":             {reflect.TypeOf(q.ALUOpAdd), constant.MakeInt64(int64(q.ALUOpAdd))},
			"ALUOpAnd":             {reflect.TypeOf(q.ALUOpAnd), constant.MakeInt64(int64(q.ALUOpAnd))},
			"ALUOpDiv":             {reflect.TypeOf(q.ALUOpDiv), constant.MakeInt64(int64(q.ALUOpDiv))},
			"ALUOpMod":             {reflect.TypeOf(q.ALUOpMod), constant.MakeInt64(int64(q.ALUOpMod))},
			"ALUOpMul":             {reflect.TypeOf(q.ALUOpMul), constant.MakeInt64(int64(q.ALUOpMul))},
			"ALUOpOr":              {reflect.TypeOf(q.ALUOpOr), constant.MakeInt64(int64(q.ALUOpOr))},
			"ALUOpShiftLeft":       {reflect.TypeOf(q.ALUOpShiftLeft), constant.MakeInt64(int64(q.ALUOpShiftLeft))},
			"ALUOpShiftRight":      {reflect.TypeOf(q.ALUOpShiftRight), constant.MakeInt64(int64(q.ALUOpShiftRight))},
			"ALUOpSub":             {reflect.TypeOf(q.ALUOpSub), constant.MakeInt64(int64(q.ALUOpSub))},
			"ALUOpXor":             {reflect.TypeOf(q.ALUOpXor), constant.MakeInt64(int64(q.ALUOpXor))},
			"ExtCPUID":             {reflect.TypeOf(q.ExtCPUID), constant.MakeInt64(int64(q.ExtCPUID))},
			"ExtInterfaceIndex":    {reflect.TypeOf(q.ExtInterfaceIndex), constant.MakeInt64(int64(q.ExtInterfaceIndex))},
			"ExtLen":               {reflect.TypeOf(q.ExtLen), constant.MakeInt64(int64(q.ExtLen))},
			"ExtLinkLayerType":     {reflect.TypeOf(q.ExtLinkLayerType), constant.MakeInt64(int64(q.ExtLinkLayerType))},
			"ExtMark":              {reflect.TypeOf(q.ExtMark), constant.MakeInt64(int64(q.ExtMark))},
			"ExtNetlinkAttr":       {reflect.TypeOf(q.ExtNetlinkAttr), constant.MakeInt64(int64(q.ExtNetlinkAttr))},
			"ExtNetlinkAttrNested": {reflect.TypeOf(q.ExtNetlinkAttrNested), constant.MakeInt64(int64(q.ExtNetlinkAttrNested))},
			"ExtPayloadOffset":     {reflect.TypeOf(q.ExtPayloadOffset), constant.MakeInt64(int64(q.ExtPayloadOffset))},
			"ExtProto":             {reflect.TypeOf(q.ExtProto), constant.MakeInt64(int64(q.ExtProto))},
			"ExtQueue":             {reflect.TypeOf(q.ExtQueue), constant.MakeInt64(int64(q.ExtQueue))},
			"ExtRXHash":            {reflect.TypeOf(q.ExtRXHash), constant.MakeInt64(int64(q.ExtRXHash))},
			"ExtRand":              {reflect.TypeOf(q.ExtRand), constant.MakeInt64(int64(q.ExtRand))},
			"ExtType":              {reflect.TypeOf(q.ExtType), constant.MakeInt64(int64(q.ExtType))},
			"ExtVLANProto":         {reflect.TypeOf(q.ExtVLANProto), constant.MakeInt64(int64(q.ExtVLANProto))},
			"ExtVLANTag":           {reflect.TypeOf(q.ExtVLANTag), constant.MakeInt64(int64(q.ExtVLANTag))},
			"ExtVLANTagPresent":    {reflect.TypeOf(q.ExtVLANTagPresent), constant.MakeInt64(int64(q.ExtVLANTagPresent))},
			"JumpBitsNotSet":       {reflect.TypeOf(q.JumpBitsNotSet), constant.MakeInt64(int64(q.JumpBitsNotSet))},
			"JumpBitsSet":          {reflect.TypeOf(q.JumpBitsSet), constant.MakeInt64(int64(q.JumpBitsSet))},
			"JumpEqual":            {reflect.TypeOf(q.JumpEqual), constant.MakeInt64(int64(q.JumpEqual))},
			"JumpGreaterOrEqual":   {reflect.TypeOf(q.JumpGreaterOrEqual), constant.MakeInt64(int64(q.JumpGreaterOrEqual))},
			"JumpGreaterThan":      {reflect.TypeOf(q.JumpGreaterThan), constant.MakeInt64(int64(q.JumpGreaterThan))},
			"JumpLessOrEqual":      {reflect.TypeOf(q.JumpLessOrEqual), constant.MakeInt64(int64(q.JumpLessOrEqual))},
			"JumpLessThan":         {reflect.TypeOf(q.JumpLessThan), constant.MakeInt64(int64(q.JumpLessThan))},
			"JumpNotEqual":         {reflect.TypeOf(q.JumpNotEqual), constant.MakeInt64(int64(q.JumpNotEqual))},
			"RegA":                 {reflect.TypeOf(q.RegA), constant.MakeInt64(int64(q.RegA))},
			"RegX":                 {reflect.TypeOf(q.RegX), constant.MakeInt64(int64(q.RegX))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
