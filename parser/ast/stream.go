package ast

import (
	"fmt"
	"tinyscript/lexer"
)

type PeekTokenStream struct {
	tokens  []*lexer.Token //TODO 保存lexer生成的tokens，更加好的方式不是全部存储，这样内存消耗会比较大
	current int
}

func NewPeekTokenStream(tokens []*lexer.Token) *PeekTokenStream {
	return &PeekTokenStream{tokens: tokens}
}

func (pt *PeekTokenStream) Next() *lexer.Token {
	if pt.current >= len(pt.tokens) {
		return nil
	}
	t := pt.tokens[pt.current]
	pt.current++
	return t
}

func (pt *PeekTokenStream) HasNext() bool {
	if pt.current >= len(pt.tokens) {
		return false
	}
	return true
}

func (pt *PeekTokenStream) Peek() *lexer.Token {
	t := pt.Next()
	if nil == t {
		return nil
	}

	pt.current -= 1
	return t
}

//参数：n 表示退回多少个token
func (pt *PeekTokenStream) PutBack(n int) {
	if pt.current-n < 0 {
		panic("putback parameter is invalid")
	}
	pt.current -= n //必须+1，因为初始化时current就指向第一个元素
}

func (pt *PeekTokenStream) NextMatch(value string) *lexer.Token {
	token := pt.Next()
	if token.Value != value {
		panic(fmt.Sprintf("syntax err: want value:%s,got %s", value, token.Value))
	}
	return token
}

func (pt *PeekTokenStream) NextMatchType(typ lexer.TokenType) *lexer.Token {
	token := pt.Next()
	if token.Typ != typ {
		panic(fmt.Sprintf("syntax err: want type: %s,got %s", token.Value, typ))
	}
	return token
}
