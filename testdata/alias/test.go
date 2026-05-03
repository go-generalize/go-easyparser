package alias

import (
	"go/types"

	"github.com/go-generalize/go-easyparser/testutil"
	tstypes "github.com/go-generalize/go-easyparser/types"
)

// Type is the expected parse result for testdata/alias/base.
var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go-easyparser/testdata/alias/base.WithMapAny": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:4:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.WithMapAny",
		Entries: map[string]tstypes.ObjectEntry{
			"extra": {
				RawName:    "Extra",
				RawTag:     `json:"extra"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:5:2"),
				Type: &tstypes.Nullable{
					Inner: &tstypes.Map{
						Key:   &tstypes.String{},
						Value: &tstypes.Any{},
					},
				},
			},
		},
	},
	"github.com/go-generalize/go-easyparser/testdata/alias/base.WithSliceAny": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:9:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.WithSliceAny",
		Entries: map[string]tstypes.ObjectEntry{
			"items": {
				RawName:    "Items",
				RawTag:     `json:"items"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:10:2"),
				Type: &tstypes.Nullable{
					Inner: &tstypes.Array{
						Inner: &tstypes.Any{},
					},
				},
			},
		},
	},
	"github.com/go-generalize/go-easyparser/testdata/alias/base.WithUserID": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:17:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.WithUserID",
		Entries: map[string]tstypes.ObjectEntry{
			"id": {
				RawName:    "ID",
				RawTag:     `json:"id"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:18:2"),
				Type:       &tstypes.String{},
			},
		},
	},
	"github.com/go-generalize/go-easyparser/testdata/alias/base.WithHandler": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:22:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.WithHandler",
		Entries: map[string]tstypes.ObjectEntry{
			"handler": {
				RawName:    "H",
				RawTag:     `json:"handler"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:23:2"),
				Type:       &tstypes.Any{},
			},
		},
	},
	"github.com/go-generalize/go-easyparser/testdata/alias/base.Box[T any]": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:27:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.Box[T any]",
		Entries: map[string]tstypes.ObjectEntry{
			"v": {
				RawName:    "V",
				RawTag:     `json:"v"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:28:2"),
				Type:       &tstypes.Any{},
			},
		},
	},
	"github.com/go-generalize/go-easyparser/testdata/alias/base.WithVecAny": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "base",
			Position: testutil.ParsePositionString("testdata/alias/base/main.go:35:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/alias/base.WithVecAny",
		Entries: map[string]tstypes.ObjectEntry{
			"items": {
				RawName:    "Items",
				RawTag:     `json:"items"`,
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/alias/base/main.go:36:2"),
				Type: &tstypes.Nullable{
					Inner: &tstypes.Array{
						Inner: &tstypes.Number{
							RawType: types.Int,
						},
					},
				},
			},
		},
	},
}
