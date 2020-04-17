package ast

var DefaultIfStmt ASTNode = MakeIfStmt()

type IfStmt struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
func MakeIfStmt() *IfStmt {
	v := &IfStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_IF_STMT)
	v.SetLabel("if_stmt")
	return v
}
