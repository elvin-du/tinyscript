package util

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
	lexer2 "tinyscript/lexer"
)

func TestNewPeekTokenStream(t *testing.T) {
	tokens := lexer2.NewLexer(bytes.NewBufferString("a+b*c"), lexer2.EndToken).Analyse()
	peekts := NewPeekTokenStream(tokens)
	assert.Equal(t, peekts.HasNext(), true)
	assertToken(t, peekts.Next(), "a", lexer2.VARIABLE)
	assertToken(t, peekts.Next(), "+", lexer2.OPERATOR)
	assertToken(t, peekts.Peek(), "b", lexer2.VARIABLE)
	assertToken(t, peekts.Next(), "b", lexer2.VARIABLE)
	peekts.PutBack(2)
	assertToken(t, peekts.Peek(), "a", lexer2.VARIABLE)
	assertToken(t, peekts.Next(), "a", lexer2.VARIABLE)
}

func assertToken(t *testing.T, token *lexer2.Token, wantValue string, wantType lexer2.TokenType) {
	assert.Equal(t, token.Typ, wantType, "err detail:"+token.String())
	assert.Equal(t, token.Value, wantValue, "err detail:"+token.String())
}
