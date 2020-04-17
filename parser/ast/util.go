package ast

import "strings"

func ToPostfixExpr(node ASTNode) string {
	if node.Type() == ASTNODE_TYPE_SCALAR || node.Type() == ASTNODE_TYPE_VARIABLE {
		return node.Lexeme().Value
	}

	arr := []string{}
	for _, child := range node.Children() {
		arr = append(arr, ToPostfixExpr(child))
	}
	str := ""
	if nil != node.Lexeme() {
		str = node.Lexeme().Value
	}

	if len(str) > 0 {
		return strings.Join(arr, " ") + " " + str
	}
	return strings.Join(arr, " ")

	//left := ""
	//right := ""
	//
	//switch node.Type() {
	//case ASTNODE_TYPE_BINARY_EXPR:
	//	left = ToPostfixExpr(node.GetChild(0))
	//	right = ToPostfixExpr(node.GetChild(1))
	//	return left + " " + right + " " + node.Lexeme().Value
	//case ASTNODE_TYPE_SCALAR, ASTNODE_TYPE_VARIABLE:
	//	return node.Lexeme().Value
	//}
	//
	//panic("ToPostfixExpr failed")
}
