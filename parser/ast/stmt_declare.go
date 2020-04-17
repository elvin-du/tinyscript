package ast

var DefaultDeclareStmt ASTNode = MakeDeclareStmt()

type DeclareStmt struct {
	*Stmt
}

func NewDeclareStmt(parent ASTNode) *DeclareStmt {
	d := MakeDeclareStmt()
	d.SetParent(parent)

	return d
}

func MakeDeclareStmt() *DeclareStmt {
	v := &DeclareStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_DECLARE_STMT)
	v.SetLabel("declare_stmt")
	return v
}

func DeclareStmtParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	stmt := NewDeclareStmt(parent)
	stream.NextMatch("var")
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
