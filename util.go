// Package easyparser provides a Go module parser for AST
package easyparser

import (
	"errors"
	"strings"

	"golang.org/x/tools/go/packages"
)

func visitErrors(pkgs []*packages.Package) error {
	errs := make([]string, 0)
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for i := range pkg.Errors {
			errs = append(errs, pkg.Errors[i].Error())
		}
	})

	if len(errs) == 0 {
		return nil
	}

	return errors.New(strings.Join(errs, "\n"))
}

// SplitPackageStruct - split package.struct
func SplitPackageStruct(s string) (string, string) {
	idx := strings.LastIndex(s, ".")

	return s[:idx], s[idx+1:]
}

// GetPackageNameFromPath - get package name from last element of path
func GetPackageNameFromPath(s string) string {
	arr := strings.Split(s, "/")

	return arr[len(arr)-1]
}
