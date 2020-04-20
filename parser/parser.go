package parser

import (
	"tinyscript/lexer"
	"tinyscript/parser/ast"
)

type Parser struct {
	stream *ast.PeekTokenStream
}

func Parse(source string) ast.ASTNode {
	return NewParser(lexer.Analyse(source)).parse()
}

func ParseFromFile(file string) ast.ASTNode {
	tokens := lexer.FromFile(file)
	return NewParser(tokens).parse()
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{stream: ast.NewPeekTokenStream(tokens)}
}

func (p *Parser) parse() ast.ASTNode {
	return ast.ProgramParse(p.stream)
}

//Expr -> digit + Expr | d|igit
//digit -> 0|1|2|....|9
func (p *Parser) SimpleParse() ast.ASTNode {
	expr := ast.MakeExpr()
	scalar := ast.NewScalar(p.stream)

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
