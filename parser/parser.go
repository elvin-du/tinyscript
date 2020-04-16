package parser

import (
	"tinyscript/lexer"
	"tinyscript/parser/ast"
	"tinyscript/parser/util"
)

type Parser struct {
	stream *util.PeekTokenStream
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{stream: util.NewPeekTokenStream(tokens)}
}

//Expr -> digit + Expr | d|igit
//digit -> 0|1|2|....|9
func (p *Parser) Parse() ast.ASTNode {
	expr := ast.NewExpr()
	scalar := ast.NewScalar(expr, p.stream)

	if !p.stream.HasNext() {
		return scalar
	}

	expr.SetLexeme(p.stream.Peek())
	p.stream.NextMatch("+")
	expr.SetLabel("+")
	expr.SetType(ast.ASTNODE_TYPE_BINARY_EXPR)
	expr.AddChild(scalar)
	rightNode := p.Parse()
	expr.AddChild(rightNode)

	return expr
}
