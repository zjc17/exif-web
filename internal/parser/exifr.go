package parser

import _ "embed"

//go:embed lite.umd.js
var liteUmdJs string

//go:embed full.umd.js
var fullUmdJs string

func ExifrScript() string {
	return liteUmdJs
}
