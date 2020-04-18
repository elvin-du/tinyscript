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

func BlockParse(stream *PeekTokenStream) ASTNode {
	stream.NextMatch("{")
	block := MakeBlock()
	for stmt := StmtParse(stream); nil != stmt; {
		block.AddChild(stmt)
		stmt = StmtParse(stream)
	}
	stream.NextMatch("}")

	return block
}
