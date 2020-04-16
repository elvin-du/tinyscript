package parser

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	lexer "tinyscript/lexer"
)

func TestParser_Parse(t *testing.T) {
	source := "1+2+3+4"
	parser := NewParser(lexer.NewLexer(bytes.NewBufferString(source), lexer.EndToken).Analyse())
	expr := parser.Parse()

	assert.Equal(t, len(expr.Children()), 2)

	v1 := expr.GetChild(0)
	assert.Equal(t, v1.Lexeme().Value, "1")
	assert.Equal(t, expr.Lexeme().Value, "+")

	e2 := expr.GetChild(1)
	v2 := e2.GetChild(0)
	assert.Equal(t, v2.Lexeme().Value, "2")
	assert.Equal(t, e2.Lexeme().Value, "+")

	e3 := e2.GetChild(1)
	v3 := e3.GetChild(0)
	assert.Equal(t, v3.Lexeme().Value, "3")
	assert.Equal(t, e3.Lexeme().Value, "+")

	v4 := e3.GetChild(1)
	assert.Equal(t, v4.Lexeme().Value, "4")
	expr.Print(0)
}
