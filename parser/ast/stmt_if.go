package ast

var DefaultIfStmt ASTNode = MakeIfStmt()

type IfStmt struct {
	*Stmt
}

func MakeIfStmt() *IfStmt {
	v := &IfStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_IF_STMT)
	v.SetLabel("if_stmt")
	return v
}

func IfStmtParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	return IfParse(parent, stream)
}

//IfStmt -> If(Expr) Block Tail
func IfParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	lexeme := stream.NextMatch("if")
	stream.NextMatch("(")
	ifStmt := MakeIfStmt()
	ifStmt.SetParent(parent)
	ifStmt.SetLexeme(lexeme)

	e := ExprParse(stream)
	ifStmt.AddChild(e)
	stream.NextMatch(")")

	block := BlockParse(parent, stream)
	ifStmt.AddChild(block)

	tail := TailParse(parent, stream)
	if tail != nil {
		ifStmt.AddChild(tail)
	}

	return ifStmt
}

//Tail -> else {Block} | else IfStmt | ⍷
func TailParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	if !stream.HasNext() || stream.Peek().Value != "else" {
		return nil
	}
	stream.NextMatch("else")
	lookahead := stream.Peek()

	if lookahead.Value == "{" {
		return BlockParse(parent, stream)
	} else if lookahead.Value == "if" {
		return IfParse(parent, stream)
	}

	return nil
}
