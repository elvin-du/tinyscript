package parser

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	lexer "tinyscript/lexer"
	"tinyscript/parser/ast"
)

func TestParser_Parse(t *testing.T) {
	source := "1+2+3+4"
	parser := NewParser(lexer.NewLexer(bytes.NewBufferString(source), lexer.EndToken).Analyse())
	expr := parser.SimpleParse()

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

func createExpr(src string) ast.ASTNode {
	tokens := lexer.NewLexer(bytes.NewBufferString(src), lexer.EndToken).Analyse()
	stream := ast.NewPeekTokenStream(tokens)
	return ast.ExprParse(stream)
}

func TestSimple(t *testing.T) {
	expr := createExpr("1+1+1")
	assert.Equal(t, ast.ToPostfixExpr(expr), "1 1 1 + +")
}

func TestSimple1(t *testing.T) {
	expr := createExpr(`"123" == ""`)
	assert.Equal(t, ast.ToPostfixExpr(expr), `"123" "" ==`)
}

func TestComplex(t *testing.T) {
	expr1 := createExpr("1+2*3")
	expr2 := createExpr("1*2+3")
	e3 := createExpr("10 * (7+4)")
	e4 := createExpr("(1*2!=7)==3!=4*5+6")

	assert.Equal(t, ast.ToPostfixExpr(expr1), "1 2 3 * +")
	assert.Equal(t, ast.ToPostfixExpr(expr2), "1 2 * 3 +")
	assert.Equal(t, ast.ToPostfixExpr(e3), "10 7 4 + *")
	assert.Equal(t, ast.ToPostfixExpr(e4), "1 2 * 7 != 3 4 5 * 6 + != ==")
}
