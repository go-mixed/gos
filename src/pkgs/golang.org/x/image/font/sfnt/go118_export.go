// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package sfnt

import (
	q "golang.org/x/image/font/sfnt"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sfnt",
		Path: "golang.org/x/image/font/sfnt",
		Deps: map[string]string{
			"errors":                             "errors",
			"fmt":                                "fmt",
			"golang.org/x/image/font":            "font",
			"golang.org/x/image/math/fixed":      "fixed",
			"golang.org/x/text/encoding/charmap": "charmap",
			"image":                              "image",
			"io":                                 "io",
			"math":                               "math",
			"sort":                               "sort",
			"strconv":                            "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Buffer":           reflect.TypeOf((*q.Buffer)(nil)).Elem(),
			"Collection":       reflect.TypeOf((*q.Collection)(nil)).Elem(),
			"Font":             reflect.TypeOf((*q.Font)(nil)).Elem(),
			"GlyphIndex":       reflect.TypeOf((*q.GlyphIndex)(nil)).Elem(),
			"LoadGlyphOptions": reflect.TypeOf((*q.LoadGlyphOptions)(nil)).Elem(),
			"NameID":           reflect.TypeOf((*q.NameID)(nil)).Elem(),
			"PostTable":        reflect.TypeOf((*q.PostTable)(nil)).Elem(),
			"Segment":          reflect.TypeOf((*q.Segment)(nil)).Elem(),
			"SegmentOp":        reflect.TypeOf((*q.SegmentOp)(nil)).Elem(),
			"Segments":         reflect.TypeOf((*q.Segments)(nil)).Elem(),
			"Units":            reflect.TypeOf((*q.Units)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrColoredGlyph": reflect.ValueOf(&q.ErrColoredGlyph),
			"ErrNotFound":     reflect.ValueOf(&q.ErrNotFound),
		},
		Funcs: map[string]reflect.Value{
			"Parse":                   reflect.ValueOf(q.Parse),
			"ParseCollection":         reflect.ValueOf(q.ParseCollection),
			"ParseCollectionReaderAt": reflect.ValueOf(q.ParseCollectionReaderAt),
			"ParseReaderAt":           reflect.ValueOf(q.ParseReaderAt),
		},
		TypedConsts: map[string]igop.TypedConst{
			"NameIDCopyright": {reflect.TypeOf(q.NameIDCopyright), constant.MakeInt64(int64(q.NameIDCopyright))},
			"SegmentOpCubeTo": {reflect.TypeOf(q.SegmentOpCubeTo), constant.MakeInt64(int64(q.SegmentOpCubeTo))},
			"SegmentOpLineTo": {reflect.TypeOf(q.SegmentOpLineTo), constant.MakeInt64(int64(q.SegmentOpLineTo))},
			"SegmentOpMoveTo": {reflect.TypeOf(q.SegmentOpMoveTo), constant.MakeInt64(int64(q.SegmentOpMoveTo))},
			"SegmentOpQuadTo": {reflect.TypeOf(q.SegmentOpQuadTo), constant.MakeInt64(int64(q.SegmentOpQuadTo))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"NameIDCompatibleFull":             {"untyped int", constant.MakeInt64(int64(q.NameIDCompatibleFull))},
			"NameIDDarkBackgroundPalette":      {"untyped int", constant.MakeInt64(int64(q.NameIDDarkBackgroundPalette))},
			"NameIDDescription":                {"untyped int", constant.MakeInt64(int64(q.NameIDDescription))},
			"NameIDDesigner":                   {"untyped int", constant.MakeInt64(int64(q.NameIDDesigner))},
			"NameIDDesignerURL":                {"untyped int", constant.MakeInt64(int64(q.NameIDDesignerURL))},
			"NameIDFamily":                     {"untyped int", constant.MakeInt64(int64(q.NameIDFamily))},
			"NameIDFull":                       {"untyped int", constant.MakeInt64(int64(q.NameIDFull))},
			"NameIDLicense":                    {"untyped int", constant.MakeInt64(int64(q.NameIDLicense))},
			"NameIDLicenseURL":                 {"untyped int", constant.MakeInt64(int64(q.NameIDLicenseURL))},
			"NameIDLightBackgroundPalette":     {"untyped int", constant.MakeInt64(int64(q.NameIDLightBackgroundPalette))},
			"NameIDManufacturer":               {"untyped int", constant.MakeInt64(int64(q.NameIDManufacturer))},
			"NameIDPostScript":                 {"untyped int", constant.MakeInt64(int64(q.NameIDPostScript))},
			"NameIDPostScriptCID":              {"untyped int", constant.MakeInt64(int64(q.NameIDPostScriptCID))},
			"NameIDSampleText":                 {"untyped int", constant.MakeInt64(int64(q.NameIDSampleText))},
			"NameIDSubfamily":                  {"untyped int", constant.MakeInt64(int64(q.NameIDSubfamily))},
			"NameIDTrademark":                  {"untyped int", constant.MakeInt64(int64(q.NameIDTrademark))},
			"NameIDTypographicFamily":          {"untyped int", constant.MakeInt64(int64(q.NameIDTypographicFamily))},
			"NameIDTypographicSubfamily":       {"untyped int", constant.MakeInt64(int64(q.NameIDTypographicSubfamily))},
			"NameIDUniqueIdentifier":           {"untyped int", constant.MakeInt64(int64(q.NameIDUniqueIdentifier))},
			"NameIDVariationsPostScriptPrefix": {"untyped int", constant.MakeInt64(int64(q.NameIDVariationsPostScriptPrefix))},
			"NameIDVendorURL":                  {"untyped int", constant.MakeInt64(int64(q.NameIDVendorURL))},
			"NameIDVersion":                    {"untyped int", constant.MakeInt64(int64(q.NameIDVersion))},
			"NameIDWWSFamily":                  {"untyped int", constant.MakeInt64(int64(q.NameIDWWSFamily))},
			"NameIDWWSSubfamily":               {"untyped int", constant.MakeInt64(int64(q.NameIDWWSSubfamily))},
		},
	})
}
