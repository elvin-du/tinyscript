package ast

var _ ASTNode = MakeIfStmt()

type IfStmt struct {
	*Stmt
}

func MakeIfStmt() *IfStmt {
	v := &IfStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_IF_STMT)
	v.SetLabel("if")
	return v
}

func IfStmtParse(stream *PeekTokenStream) ASTNode {
	return IfParse(stream)
}

//IfStmt -> If(Expr) Block Tail
func IfParse(stream *PeekTokenStream) ASTNode {
	lexeme := stream.NextMatch("if")
	stream.NextMatch("(")
	ifStmt := MakeIfStmt()
	ifStmt.SetLexeme(lexeme)

	e := ExprParse(stream)
	ifStmt.AddChild(e)
	stream.NextMatch(")")

	block := BlockParse(stream)
	ifStmt.AddChild(block)

	tail := TailParse(stream)
	if tail != nil {
		ifStmt.AddChild(tail)
	}

	return ifStmt
}

//Tail -> else {Block} | else IfStmt | ⍷
func TailParse(stream *PeekTokenStream) ASTNode {
	if !stream.HasNext() || stream.Peek().Value != "else" {
		return nil
	}
	stream.NextMatch("else")
	lookahead := stream.Peek()

	if lookahead.Value == "{" {
		return BlockParse(stream)
	} else if lookahead.Value == "if" {
		return IfParse(stream)
	}

	return nil
}

func (i *IfStmt) GetExpr() ASTNode {
	return i.GetChild(0)
}

func (i *IfStmt) GetBlock() ASTNode {
	return i.GetChild(1)
}
func (i *IfStmt) GetElseBlock() ASTNode {
	block := i.GetChild(2)
	if block != nil && block.Type() == ASTNODE_TYPE_BLOCK {
		return block
	}

	return nil
}
func (i *IfStmt) GetElseIfStmt() ASTNode {
	ifStmt := i.GetChild(2)
	if ifStmt != nil && ifStmt.Type() == ASTNODE_TYPE_IF_STMT {
		return ifStmt
	}
	return nil
}
