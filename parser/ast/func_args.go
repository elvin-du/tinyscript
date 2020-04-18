package ast

var _ ASTNode = &Factor{}

type FuncArgs struct {
	*node
}

func MakeFuncArgs() *FuncArgs {
	s := &FuncArgs{MakeNode()}
	s.SetLabel("args")
	return s
}

func FuncArgsParse(stream *PeekTokenStream) ASTNode {
	args := MakeFuncArgs()
	for ; stream.Peek().IsType(); {
		typ := stream.Next()
		v := FactorParse(stream)
		v.SetTypeLexeme(typ)
		args.AddChild(v)
		if stream.Peek().Value != ")" {
			stream.NextMatch(",")
		}
	}

	return args
}
