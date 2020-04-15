package lexer

import "fmt"

type TokenType int

const (
	KEYWORD  TokenType = 1
	VARIABLE TokenType = 2
	OPERATOR TokenType = 3
	BRACKET  TokenType = 4
	STRING   TokenType = 5
	FLOAT    TokenType = 6
	BOOLEAN  TokenType = 7
	INTEGER  TokenType = 8
)

func (tt TokenType) String() string {
	switch tt {

	case KEYWORD:
		return "keyword"
	case VARIABLE:
		return "variable"
	case OPERATOR:
		return "operator"
	case BRACKET:
		return "bracket"
	case STRING:
		return "string`"
	case FLOAT:
		return "float"
	case BOOLEAN:
		return "boolean"
	case INTEGER:
		return "integer"
	}

	panic("unexpected token type")
}

type Token struct {
	Typ   TokenType
	Value string
}

func NewToken(t TokenType, v string) *Token {
	return &Token{Typ: t, Value: v}
}

func (t *Token) IsVariable() bool {
	return t.Typ == VARIABLE
}

func (t *Token) IsScalar() bool {
	return t.Typ == FLOAT || t.Typ == BOOLEAN || t.Typ == INTEGER || t.Typ == STRING
}

func (t *Token) IsNumber() bool {
	return t.Typ == INTEGER || t.Typ == FLOAT
}

func (t *Token) IsOperator() bool {
	return t.Typ == OPERATOR
}

func (t *Token) String() string {
	return fmt.Sprintf("type:%v,value:%s", t.Typ, t.Value)
}
