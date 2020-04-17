package ast

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/lexer"
)

func TestAssignStmtParse(t *testing.T) {
	src := "i = 100*2"
	tokens := lexer.NewLexer(bytes.NewBufferString(src), lexer.EndToken).Analyse()
	stream := NewPeekTokenStream(tokens)
	stmt := AssignStmtParse(nil, stream)
	assert.Equal(t, ToPostfixExpr(stmt), "i 100 2 * =")
}
