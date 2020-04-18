package ast

var _ ASTNode = MakeForStmt()

type ForStmt struct {
	*Stmt
}

func MakeForStmt() *ForStmt {
	v := &ForStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_FOR_STMT)
	v.SetLabel("for")
	return v
}
