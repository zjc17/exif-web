package parser

import _ "embed"

//go:embed lite.umd.js
var liteUmdJs string

func ExifrScript() string {
	return liteUmdJs
}
