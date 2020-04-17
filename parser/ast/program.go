package ast

var DefaultProgram ASTNode = &Block{}

type Program struct {
	*Block
}

func MakeProgram() *Program {
	b := &Program{MakeBlock()}
	b.SetLabel("program")
	return b
}

func ProgramParse(parent ASTNode, stream *PeekTokenStream) ASTNode {
	p := MakeProgram()
	p.SetParent(parent)
	for stmt := StmtParse(parent, stream); nil != stmt; {
		p.AddChild(stmt)
		stmt = StmtParse(parent, stream)
	}

	return p
}
