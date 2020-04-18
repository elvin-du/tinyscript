package ast

var _ ASTNode = &ReturnStmt{}

type ReturnStmt struct {
	*Stmt
}

func MakeReturnStmt() *ReturnStmt {
	v := &ReturnStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_RETURN_STMT)
	v.SetLabel("return")
	return v
}

func ReturnStmtParse(stream *PeekTokenStream) ASTNode {
	var lexeme = stream.NextMatch("return")
	var expr = ExprParse(stream)

	var stmt = MakeReturnStmt()
	stmt.SetLexeme(lexeme)
	if expr != nil {
		stmt.AddChild(expr)
	}

	return stmt
}
