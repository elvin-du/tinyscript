package ast

var DefaultAssignStmt ASTNode = MakeAssignStmt()

type AssignStmt struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
func MakeAssignStmt() *AssignStmt {
	v := &AssignStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_ASSIGN_STMT)
	v.SetLabel("assign_stmt")
	return v
}
