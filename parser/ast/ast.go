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
	TypeLexeme() *lexer.Token
	IsValueType() bool
	Prop(string) interface{}

	//set
	AddChild(ASTNode)
	SetLexeme(*lexer.Token)
	SetTypeLexeme(*lexer.Token)
	SetType(NodeType)
	SetLabel(string)
	SetParent(ASTNode)
	SetProp(string, interface{})
}

type node struct {
	parent     ASTNode
	children   []ASTNode
	label      string
	typ        NodeType
	lexeme     *lexer.Token
	typeLexeme *lexer.Token
	prop       map[string]interface{}
}

//test
var _ ASTNode = &node{}

func MakeNode() *node {
	return &node{children: make([]ASTNode, 0), prop: make(map[string]interface{})}
}
func (n *node) Prop(key string) interface{} {
	return n.prop[key]
}
func (n *node) SetProp(key string, value interface{}) {
	n.prop[key] = value
}
func (n *node) Lexeme() *lexer.Token {
	return n.lexeme
}
func (n *node) TypeLexeme() *lexer.Token {
	return n.typeLexeme
}
func (n *node) IsValueType() bool {
	return n.typ == ASTNODE_TYPE_VARIABLE || n.typ == ASTNODE_TYPE_SCALAR
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
	if int(index) >= len(n.children) {
		return nil
	}
	return n.children[index]
}
func (n *node) Parent() ASTNode {
	return n.parent
}
func (n *node) AddChild(node ASTNode) {
	node.SetParent(n)
	n.children = append(n.children, node)
}
func (n *node) SetLexeme(lexeme *lexer.Token) {
	n.lexeme = lexeme
}
func (n *node) SetTypeLexeme(lexeme *lexer.Token) {
	n.typeLexeme = lexeme
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
