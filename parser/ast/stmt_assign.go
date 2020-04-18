package ast

var DefaultAssignStmt ASTNode = MakeAssignStmt()

type AssignStmt struct {
	*Stmt
}

func MakeAssignStmt() *AssignStmt {
	v := &AssignStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_ASSIGN_STMT)
	v.SetLabel("assign_stmt")
	return v
}

func AssignStmtParse(stream *PeekTokenStream) ASTNode {
	stmt := MakeAssignStmt()
	//stmt.SetParent(parent)
	tkn := stream.Peek()
	factor := FactorParse(stream)
	if nil == factor {
		panic("syntax error:" + tkn.String())
	}
	stmt.AddChild(factor)
	lexeme := stream.NextMatch("=")
	stmt.SetLexeme(lexeme)
	expr := ExprParse(stream)
	stmt.AddChild(expr)

	return stmt
}
