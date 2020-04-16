package ast

import (
	"fmt"
	"strings"
	"tinyscript/lexer"
)

type ASTNode interface {
	//get
	Lexeme() *lexer.Token
	Type() NodeType
	Label() string
	Children() []ASTNode
	GetChild(uint) ASTNode
	Parent() ASTNode
	Print(indent int)

	//set
	AddChild(ASTNode)
	SetLexeme(*lexer.Token)
	SetType(NodeType)
	SetLabel(string)
	SetParent(ASTNode)
}

type node struct {
	parent   ASTNode
	children []ASTNode
	label    string
	typ      NodeType
	lexeme   *lexer.Token
}

//test
var _ ASTNode = &node{}

func NewNode() *node {
	return &node{children: make([]ASTNode, 0)}
}
func (n *node) Lexeme() *lexer.Token {
	return n.lexeme
}
func (n *node) Type() NodeType {
	return n.typ
}
func (n *node) Label() string {
	return n.label
}
func (n *node) Children() []ASTNode {
	return n.children
}
func (n *node) GetChild(index uint) ASTNode {
	return n.children[index]
}
func (n *node) Parent() ASTNode {
	return n.parent
}
func (n *node) AddChild(node ASTNode) {
	n.children = append(n.children, node)
}
func (n *node) SetLexeme(lexeme *lexer.Token) {
	n.lexeme = lexeme
}
func (n *node) SetType(t NodeType) {
	n.typ = t
}
func (n *node) SetLabel(str string) {
	n.label = str
}
func (n *node) SetParent(node ASTNode) {
	n.parent = node
}
func (n *node) Print(indent int) {
	fmt.Printf("%s%s\n", strings.Repeat("  ", indent*2), n.label)
	for _, child := range n.children {
		child.Print(indent + 2)
	}
}
