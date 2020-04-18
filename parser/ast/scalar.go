package ast

var _ ASTNode = &Factor{}

type Scalar struct {
	*Factor
}

func NewScalar(stream *PeekTokenStream) *Scalar {
	return &Scalar{NewFactor(stream)}
}

func MakeScalar() *Scalar {
	s := &Scalar{MakeFactor()}
	s.SetType(ASTNODE_TYPE_SCALAR)
	return s
}
