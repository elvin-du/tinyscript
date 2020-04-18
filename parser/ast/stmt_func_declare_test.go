package ast

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/lexer"
)

func TestFuncDeclareStmtParse(t *testing.T) {
	stream := NewPeekTokenStream(lexer.FromFile("./../../tests/function.ts"))
	stmt := StmtParse(stream).(*FuncDeclareStmt)
	args := stmt.Args()
	assert.Equal(t, args.GetChild(0).Lexeme().Value, "a")
	assert.Equal(t, args.GetChild(1).Lexeme().Value, "b")

	typ := stmt.FuncType()
	assert.Equal(t, typ, "int")

	funcVariable := stmt.FuncVariable()
	assert.Equal(t, funcVariable.Lexeme().Value, "add")

	block := stmt.Block()
	assert.Equal(t, block.GetChild(0).Lexeme().Value, "return")
}

func TestFunctionRecursion(t *testing.T) {
	stream := NewPeekTokenStream(lexer.FromFile("./../../tests/recursion.ts"))
	stmt := StmtParse(stream).(*FuncDeclareStmt)
	assert.Equal(t, ToBFSString(stmt, 4), "func fact args block")
	assert.Equal(t, ToBFSString(stmt.Args(), 2), "args n")
	assert.Equal(t, ToBFSString(stmt.Block(), 3), "block if return")
}
