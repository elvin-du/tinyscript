package ast

import (
	"tinyscript/lexer"
	"tinyscript/parser/util"
)

var _ ASTNode = &Factor{}

type Factor struct {
	*node
}

func NewFactor(parent ASTNode, stream *util.PeekTokenStream) *Factor {
	factor := &Factor{NewNode()}
	token := stream.Next()
	factor.SetLexeme(token)
	factor.SetLabel(token.Value)
	factor.SetParent(parent)

	if lexer.VARIABLE == token.Typ {
		factor.SetType(ASTNODE_TYPE_VARIABLE)
	} else {
		factor.SetType(ASTNODE_TYPE_SCALAR)
	}

	return factor
}
