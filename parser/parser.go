package parser

import (
	"tinyscript/lexer"
	"tinyscript/parser/ast"
)

type Parser struct {
	stream *ast.PeekTokenStream
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{stream: ast.NewPeekTokenStream(tokens)}
}

//Expr -> digit + Expr | d|igit
//digit -> 0|1|2|....|9
func (p *Parser) SimpleParse() ast.ASTNode {
	expr := ast.MakeExpr()
	scalar := ast.NewScalar(expr, p.stream)

	if !p.stream.HasNext() {
		return scalar
	}

	expr.SetLexeme(p.stream.Peek())
	p.stream.NextMatch("+")
	expr.SetLabel("+")
	expr.SetType(ast.ASTNODE_TYPE_BINARY_EXPR)
	expr.AddChild(scalar)
	rightNode := p.SimpleParse()
	expr.AddChild(rightNode)

	return expr
}
