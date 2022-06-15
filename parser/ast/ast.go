package ast

import (
	"fmt"
	"strings"
	"tinyscript/lexer"
)

type ASTNode interface {
	//get
	Lexeme() *lexer.Token //ast节点对应的token是什么
	Type() NodeType
	Label() string //用字符串标识ast节点的含义，主要用于打印日志
	Children() []ASTNode
	GetChild(uint) ASTNode
	Parent() ASTNode
	Print(indent int)
	TypeLexeme() *lexer.Token //标识变量的类型和函数的返回值类型；别的ast节点没有必要设置这个属性
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
	label      string //备注（标签）
	typ        NodeType
	lexeme     *lexer.Token           //词法单元
	typeLexeme *lexer.Token           //func foo(int a); 这时typelexeme等于int型的token
	prop       map[string]interface{} //用于符号表，语法分析不会用到
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
