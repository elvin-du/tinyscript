package ast

var DefaultProgram ASTNode = &Block{}

type Program struct {
	*Block
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
//func MakeVariable() *Variable {
//	v := &Variable{MakeFactor()}
//	v.SetType(ASTNODE_TYPE_VARIABLE)
//	return v
//}
