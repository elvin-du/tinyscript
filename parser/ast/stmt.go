package ast

var DefaultStmt ASTNode = MakeStmt()

type Stmt struct {
	*node
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
func MakeStmt() *Stmt {
	s := &Stmt{MakeNode()}
	return s
}
