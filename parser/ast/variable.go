package ast

var _ ASTNode = &Variable{}

type Variable struct {
	*Factor
}

func NewVariable(stream *PeekTokenStream) *Variable {
	return &Variable{NewFactor(stream)}
}

func MakeVariable() *Variable {
	v := &Variable{MakeFactor()}
	v.SetType(ASTNODE_TYPE_VARIABLE)
	return v
}
