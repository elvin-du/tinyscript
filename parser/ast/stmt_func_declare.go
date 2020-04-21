package ast

var _ ASTNode = MakeFuncDeclareStmt()

type FuncDeclareStmt struct {
	*Stmt
}

func MakeFuncDeclareStmt() *FuncDeclareStmt {
	v := &FuncDeclareStmt{MakeStmt()}
	v.SetType(ASTNODE_TYPE_FUNCTION_DECLARE_STMT)
	v.SetLabel("func")
	return v
}

func FuncDeclareStmtParse(stream *PeekTokenStream) *FuncDeclareStmt {
	stream.NextMatch("func")
	//func add() int {}
	fn := MakeFuncDeclareStmt()
	lexeme := stream.Peek()
	fnV := FactorParse(stream)
	fn.SetLexeme(lexeme)
	fn.AddChild(fnV)

	stream.NextMatch("(")
	args := FuncArgsParse(stream)
	stream.NextMatch(")")
	fn.AddChild(args)

	keyword := stream.Next()
	if !keyword.IsType() {
		panic("syntax error: unexpected " + keyword.Value)
	}

	fnV.SetTypeLexeme(keyword)
	block := BlockParse(stream)
	fn.AddChild(block)

	return fn
}

func (f *FuncDeclareStmt) FuncVariable() ASTNode {
	return f.GetChild(0)
}
func (f *FuncDeclareStmt) Args() ASTNode {
	return f.GetChild(1)
}
func (f *FuncDeclareStmt) FuncType() string {
	return f.FuncVariable().TypeLexeme().Value
}
func (f *FuncDeclareStmt) Block() ASTNode {
	return f.GetChild(2)
}
