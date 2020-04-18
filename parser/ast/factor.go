package ast

import (
	"tinyscript/lexer"
)

var _ ASTNode = &Factor{}

type Factor struct {
	*node
}

func MakeFactor() *Factor {
	return &Factor{MakeNode()}
}

func NewFactor(stream *PeekTokenStream) *Factor {
	factor := &Factor{MakeNode()}
	token := stream.Next()
	factor.SetLexeme(token)
	factor.SetLabel(token.Value)

	if lexer.VARIABLE == token.Typ {
		factor.SetType(ASTNODE_TYPE_VARIABLE)
	} else {
		factor.SetType(ASTNODE_TYPE_SCALAR)
	}

	return factor
}

func FactorParse(stream *PeekTokenStream) ASTNode {
	token := stream.Peek()
	typ := token.Typ
	if lexer.VARIABLE == typ {
		stream.Next()
		v := MakeVariable()
		v.SetLabel(token.Value)
		v.SetLexeme(token)
		return v
	} else if token.IsScalar() {
		stream.Next()
		scalar := MakeScalar()
		scalar.SetLabel(token.Value)
		scalar.SetLexeme(token)
		return scalar
	}
	return nil
}
