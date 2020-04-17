package ast

var DefaultBlock ASTNode = MakeBlock()

type Block struct {
	*Stmt
}

//func NewVariable(parent ASTNode, stream *PeekTokenStream) *Variable {
//	return &Variable{NewFactor(parent, stream)}
//}
//
func MakeBlock() *Block {
	b := &Block{MakeStmt()}
	b.SetType(ASTNODE_TYPE_BLOCK)
	b.SetLabel("block")
	return b
}
