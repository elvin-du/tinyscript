package ast

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/lexer"
)

func TestDeclareStmtParse(t *testing.T) {
	src := "var i = 100*2"
	tokens := lexer.NewLexer(bytes.NewBufferString(src), lexer.EndToken).Analyse()
	stream := NewPeekTokenStream(tokens)
	stmt := DeclareStmtParse(stream)
	assert.Equal(t, ToPostfixExpr(stmt), "i 100 2 * =")
}
