package ast

var DefaultFuncDeclareStmt ASTNode = MakeFuncDeclareStmt()

type FuncDeclareStmt struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
func MakeFuncDeclareStmt() *FuncDeclareStmt {
	v := &FuncDeclareStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_DECLARE_STMT)
	v.SetLabel("func_declare_stmt")
	return v
}
