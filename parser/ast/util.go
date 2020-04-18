package ast

import (
	"container/list"
	"strings"
)

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

func ToBFSString(node ASTNode, max int) string {
	l := list.New()
	l.PushBack(node)
	strs := []string{}
	for e, i := l.Front(), 0; nil != e && i < max; e = l.Front() {
		i += 1
		parent := l.Remove(e).(ASTNode)
		strs = append(strs, parent.Label())

		for _, child := range parent.Children() {
			l.PushBack(child)
		}
	}

	return strings.Join(strs, " ")
}
