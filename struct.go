// Package easyparser provides a Go module parser for AST
package easyparser

import (
	"go/types"
	"reflect"
	"sort"
	"strings"

	tstypes "github.com/go-generalize/go-easyparser/types"
)

const (
	jsonTagOmitempty = "omitempty"
)

func (p *pkgParser) parseStruct(strct *types.Struct) tstypes.Type {
	type entryPair struct {
		key   string
		value tstypes.ObjectEntry
	}

	entries := make([][]entryPair, strct.NumFields())
	// embedding
	for i := 0; i < strct.NumFields(); i++ {
		v := strct.Field(i)
		tag := strct.Tag(i)

		if !v.Exported() || !v.Embedded() {
			continue
		}

		jsonTag := strings.SplitN(reflect.StructTag(tag).Get("json"), ",", 2)
		field := ""
		if len(jsonTag) >= 1 {
			field = jsonTag[0]
		}
		optional := len(jsonTag) >= 2 && jsonTag[1] == jsonTagOmitempty

		if field == "-" && !p.IgnoreOmittedJSONField {
			continue
		}

		tst := p.parseType(v.Type(), true)
		if len(field) == 0 {
			if o, ok := tst.(*tstypes.Object); ok {
				fieldEntries := make([]entryPair, 0, len(o.Entries))
				for k, v := range o.Entries {
					fieldEntries = append(fieldEntries, entryPair{k, v})
				}
				sort.Slice(fieldEntries, func(i, j int) bool {
					return fieldEntries[i].value.FieldIndex < fieldEntries[j].value.FieldIndex
				})
				entries[i] = fieldEntries

				continue
			}

			field = v.Name()
		}

		entries[i] = []entryPair{
			{
				key: field,
				value: tstypes.ObjectEntry{
					RawName:    v.Name(),
					RawTag:     tag,
					Type:       tst,
					Optional:   optional,
					FieldIndex: i,
				},
			},
		}
	}

	// not embedding
	for i := 0; i < strct.NumFields(); i++ {
		v := strct.Field(i)
		tag := strct.Tag(i)

		if !v.Exported() || v.Embedded() {
			continue
		}

		obj, _, _ := types.LookupFieldOrMethod(strct, true, p.pkg, v.Name())
		pos := p.fset.Position(obj.Pos())

		jsonTag := strings.SplitN(reflect.StructTag(tag).Get("json"), ",", 2)
		field := ""
		if len(jsonTag) >= 1 {
			field = jsonTag[0]
		}
		optional := len(jsonTag) >= 2 && jsonTag[1] == jsonTagOmitempty

		if field == "-" && !p.IgnoreOmittedJSONField {
			continue
		}

		if len(field) == 0 {
			field = v.Name()
		}

		tst := p.parseType(v.Type(), true)
		if optional {
			entries[i] = []entryPair{
				{
					key: field,
					value: tstypes.ObjectEntry{
						RawName:  v.Name(),
						RawTag:   tag,
						Type:     p.removeNullable(tst),
						Optional: true,
						Position: &pos,
					},
				},
			}
		} else {
			entries[i] = []entryPair{
				{
					key: field,
					value: tstypes.ObjectEntry{
						RawName:  v.Name(),
						RawTag:   tag,
						Type:     tst,
						Optional: false,
						Position: &pos,
					},
				},
			}
		}
	}

	obj := tstypes.Object{
		Entries: map[string]tstypes.ObjectEntry{},
	}

	idx := 0
	for _, entry := range entries {
		for _, e := range entry {
			e.value.FieldIndex = idx
			obj.Entries[e.key] = e.value
			idx++
		}
	}

	return &obj
}

func (p *Parser) removeNullable(typ tstypes.Type) tstypes.Type {
	if nullable, ok := typ.(*tstypes.Nullable); ok {
		return nullable.Inner
	}

	return typ
}
