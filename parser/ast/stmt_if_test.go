package ast

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/lexer"
)

func TestIfStmtParse(t *testing.T) {
	stream := createTokenStream(`if(a){
a = 1
}`)
	stmt := IfStmtParse(nil, stream)
	e := stmt.GetChild(0)
	block := stmt.GetChild(1)
	assignStmt := block.GetChild(0)

	assert.Equal(t, e.Lexeme().Value, "a")
	assert.Equal(t, assignStmt.Lexeme().Value, "=")
}

func createTokenStream(src string) *PeekTokenStream {
	tokens := lexer.NewLexer(bytes.NewBufferString(src), lexer.EndToken).Analyse()
	stream := NewPeekTokenStream(tokens)
	return stream
}

func TestIfElseStmtParse(t *testing.T) {
	stream := createTokenStream(`if(a){
a = 1
}else{
a = 2
a = a * 3
}`)

	stmt := IfStmtParse(nil, stream)
	expr := stmt.GetChild(0)
	block := stmt.GetChild(1)
	assignStmt := block.GetChild(0)
	elseBlock := stmt.GetChild(2)
	assignStmt2 := elseBlock.GetChild(0)

	assert.Equal(t, expr.Lexeme().Value, "a")
	assert.Equal(t, assignStmt.Lexeme().Value, "=")
	assert.Equal(t, assignStmt2.Lexeme().Value, "=")
	assert.Equal(t, len(elseBlock.Children()), 2)
}
