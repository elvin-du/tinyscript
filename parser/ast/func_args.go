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
	for stream.Peek().IsType() {
		typ := stream.Next()
		v := FactorParse(stream)
		v.SetTypeLexeme(typ) //为语义分析做准备，设置参数变量的类型
		args.AddChild(v)
		if stream.Peek().Value != ")" {
			stream.NextMatch(",")
		}
	}

	return args
}
