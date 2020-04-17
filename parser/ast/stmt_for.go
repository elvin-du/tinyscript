package ast

var DefaultForStmt ASTNode = MakeForStmt()

type ForStmt struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
func MakeForStmt() *ForStmt {
	v := &ForStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_FOR_STMT)
	v.SetLabel("for_stmt")
	return v
}
