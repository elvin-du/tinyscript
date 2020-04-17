package ast

func ToPostfixExpr(node ASTNode) string {
	left := ""
	right := ""

	switch node.Type() {
	case ASTNODE_TYPE_BINARY_EXPR:
		left = ToPostfixExpr(node.GetChild(0))
		right = ToPostfixExpr(node.GetChild(1))
		return left + " " + right + " " + node.Lexeme().Value
	case ASTNODE_TYPE_SCALAR, ASTNODE_TYPE_VARIABLE:
		return node.Lexeme().Value
	}

	panic("ToPostfixExpr failed")
}
