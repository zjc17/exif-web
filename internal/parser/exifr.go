package parser

import (
	_ "embed"
)

//go:embed full.umd.js
var fullUmdJs string

func ExifrScript() string {
	return fullUmdJs
}
