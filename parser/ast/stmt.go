package ast

var DefaultStmt ASTNode = MakeStmt()

type Stmt struct {
	*node
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
func MakeStmt() *Stmt {
	s := &Stmt{MakeNode()}
	return s
}

func StmtParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	token := stream.Next()
	lookahead := stream.Peek()
	stream.PutBack(1)

	if token.IsVariable() && lookahead != nil && lookahead.Value == "=" {
		return AssignStmtParse(parent, stream)
	} else if token.Value == "var" {
		return DeclareStmtParse(parent, stream)
	}

	return nil
}
