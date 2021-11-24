package replace

import (
	"github.com/go-generalize/go-easyparser/testutil"
	tstypes "github.com/go-generalize/go-easyparser/types"
)

var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go-easyparser/testdata/replace/base.Struct": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/replace/base/main.go:6:6"),
		},
		Name: "github.com/go-generalize/go-easyparser/testdata/replace/base.Struct",
		Entries: map[string]tstypes.ObjectEntry{
			"A": {
				RawName:    "A",
				FieldIndex: 0,

				Type:     &tstypes.Number{},
				Optional: false,
				Position: testutil.ParsePositionString("testdata/replace/base/main.go:7:2"),
			},
		},
	},
}
