package ast

var DefaultDeclareStmt ASTNode = MakeDeclareStmt()

type DeclareStmt struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
func MakeDeclareStmt() *DeclareStmt {
	v := &DeclareStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_DECLARE_STMT)
	v.SetLabel("declare_stmt")
	return v
}
