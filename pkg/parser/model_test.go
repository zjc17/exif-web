package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseResult_From_Error(t *testing.T) {
	p := new(ParseResult)
	err := p.From("")
	assert.NotNil(t, err)
}
