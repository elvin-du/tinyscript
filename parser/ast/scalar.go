package ast

import "tinyscript/parser/util"

var _ ASTNode = &Factor{}

type Scalar struct {
	*Factor
}

func NewScalar(parent ASTNode, stream *util.PeekTokenStream) *Scalar {
	return &Scalar{NewFactor(parent, stream)}
}
