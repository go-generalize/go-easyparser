package main

import (
	"github.com/go-generalize/go-easyparser/testdata/conflict/base/pkg"
)

type Data struct {
	Hoge    Hoge
	PkgHoge pkg.Hoge
}

type Hoge struct {
	Data int
}
