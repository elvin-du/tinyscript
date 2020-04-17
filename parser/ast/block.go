package ast

var DefaultBlock ASTNode = MakeBlock()

type Block struct {
	*Stmt
}

func MakeBlock() *Block {
	b := &Block{MakeStmt()}
	b.SetType(ASTNODE_TYPE_BLOCK)
	b.SetLabel("block")
	return b
}

func BlockParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	stream.NextMatch("{")
	block := MakeBlock()
	block.SetParent(parent)
	for stmt := StmtParse(parent, stream); nil != stmt; {
		block.AddChild(stmt)
		stmt = StmtParse(parent, stream)
	}
	stream.NextMatch("}")

	return block
}
