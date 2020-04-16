package ast

import "tinyscript/parser/util"

var _ ASTNode = &Variable{}

type Variable struct {
	*Factor
}

func NewVariable(parent ASTNode, stream *util.PeekTokenStream) *Variable {
	return &Variable{NewFactor(parent, stream)}
}
