package ast

var DefaultStmt ASTNode = MakeStmt()

type Stmt struct {
	*node
}

func MakeStmt() *Stmt {
	s := &Stmt{MakeNode()}
	return s
}

func StmtParse(stream *PeekTokenStream) ASTNode {
	token := stream.Next()
	lookahead := stream.Peek()
	stream.PutBack(1)

	if token.IsVariable() && lookahead != nil && lookahead.Value == "=" {
		return AssignStmtParse(stream)
	} else if token.Value == "var" {
		return DeclareStmtParse(stream)
	} else if token.Value == "func" {
		return FuncDeclareStmtParse(stream)
	} else if token.Value == "return" {
		return ReturnStmtParse(stream)
	} else if token.Value == "if" {
		return IfStmtParse(stream)
	}

	return nil
}
