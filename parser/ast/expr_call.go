package ast

var _ ASTNode = &CallExpr{}

type CallExpr struct {
	*node
}

func MakeCallExpr() *CallExpr {
	e := &CallExpr{MakeNode()}
	e.SetType(ASTNODE_TYPE_CALL_EXPR)
	return e
}

func CallExprParse(factor ASTNode, stream *PeekTokenStream) ASTNode {
	expr := MakeCallExpr()
	expr.AddChild(factor)
	stream.NextMatch("(")
	p := ExprParse(stream)
	for ; p != nil; p = ExprParse(stream) {
		expr.AddChild(p)
		if stream.Peek().Value != ")" {
			stream.NextMatch(",")
		}
	}

	stream.NextMatch(")")
	return expr
}
