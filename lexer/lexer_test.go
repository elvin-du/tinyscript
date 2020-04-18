package lexer

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"regexp"
	"strings"
	"testing"
)

func TestLexer_MakeVarOrKeyword(t *testing.T) {
	l := NewLexer(bytes.NewBufferString("if abc"), "$")
	token := l.MakeVarOrKeyword()
	token2 := NewLexer(bytes.NewBufferString("true abc"), EndToken).MakeVarOrKeyword()

	assert.Equal(t, token.Typ, KEYWORD)
	assert.Equal(t, token.Value, "if")
	assert.Equal(t, token2.Typ, BOOLEAN)
	assert.Equal(t, token2.Value, "true")

	l.Next()
	token3 := l.MakeVarOrKeyword()
	assert.Equal(t, token3.Typ, VARIABLE)
	assert.Equal(t, token3.Value, "abc")
}

func TestLexer_MakeString(t *testing.T) {
	token := NewLexer(bytes.NewBufferString(`"123"`), "$").MakeString()
	assert.Equal(t, token.Typ, STRING)
	assert.Equal(t, token.Value, `"123"`)
}

func TestLexer_MakeOp(t *testing.T) {
	tests := []string{
		"+ xxx",
		"++mmm",
		"/=g",
		"==1",
		"&=3434",
		"&8888",
		"||xxxx",
		"^=111",
		"%79",
	}

	for _, test := range tests {
		token := NewLexer(bytes.NewBufferString(test), "$").MakeOp()
		assert.Equal(t, token.Typ, OPERATOR)
	}
}
func TestLexer_MakeNumber(t *testing.T) {
	tests := []string{
		"+0 aa",
		"-0 aa",
		".3000 aa",
		".55 ww",
		"778.99 aa",
		"355 kkk",
		"-888*234aa",
	}

	for _, test := range tests {
		token := NewLexer(bytes.NewBufferString(test), "$").MakeNumber()

		value := regexp.MustCompile("[* ]+").Split(token.Value, 1)[0]
		//t.Log(value)
		assert.Equal(t, token.Value, value)
		if strings.Contains(token.Value, ".") {
			assert.Equal(t, token.Typ, FLOAT)
		} else {
			assert.Equal(t, token.Typ, INTEGER)
		}
	}
}

func TestLexer_Analyse(t *testing.T) {
	source := `(w+c)^100.12==+100-30eee`
	lexer := NewLexer(bytes.NewBufferString(source), EndToken)
	tokens := lexer.Analyse()
	assert.Equal(t, len(tokens), 12)
	assert.Equal(t, tokens[0].Value, "(")
	assert.Equal(t, tokens[1].Value, "w")
	assert.Equal(t, tokens[2].Value, "+")
	assert.Equal(t, tokens[3].Value, "c")
	assert.Equal(t, tokens[4].Value, ")")
	assert.Equal(t, tokens[5].Value, "^")
	assert.Equal(t, tokens[6].Value, "100.12")
	assert.Equal(t, tokens[7].Value, "==")
	assert.Equal(t, tokens[8].Value, "+100")
	assert.Equal(t, tokens[9].Value, "-")
	assert.Equal(t, tokens[10].Value, "30")
	assert.Equal(t, tokens[11].Value, "eee")
}

func Test_Function(t *testing.T) {
	source := `func foo(a,b){ 
		print(a+b) 
		}
		foo(-100.0,100)`

	lexer := NewLexer(bytes.NewBufferString(source), EndToken)
	tokens := lexer.Analyse()

	assertToken(t, tokens[0], "func", KEYWORD)
	assertToken(t, tokens[1], "foo", VARIABLE)
	assertToken(t, tokens[2], "(", BRACKET)
	assertToken(t, tokens[3], "a", VARIABLE)
	assertToken(t, tokens[4], ",", OPERATOR)
	assertToken(t, tokens[5], "b", VARIABLE)
	assertToken(t, tokens[6], ")", BRACKET)
	assertToken(t, tokens[7], "{", BRACKET)
	assertToken(t, tokens[8], "print", VARIABLE)
	assertToken(t, tokens[9], "(", BRACKET)
	assertToken(t, tokens[10], "a", VARIABLE)
	assertToken(t, tokens[11], "+", OPERATOR)
	assertToken(t, tokens[12], "b", VARIABLE)
	assertToken(t, tokens[13], ")", BRACKET)
	assertToken(t, tokens[14], "}", BRACKET)
	assertToken(t, tokens[15], "foo", VARIABLE)
	assertToken(t, tokens[16], "(", BRACKET)
	assertToken(t, tokens[17], "-100.0", FLOAT)
	assertToken(t, tokens[18], ",", OPERATOR)
	assertToken(t, tokens[19], "100", INTEGER)
	assertToken(t, tokens[20], ")", BRACKET)
}

func TestDeleteComment(t *testing.T) {
	source := `/*12324abdfda
				34fa9kfjl*/a=1
		`
	lexer := NewLexer(bytes.NewBufferString(source), EndToken)
	tokens := lexer.Analyse()
	assert.Equal(t, len(tokens), 3)
}

func assertToken(t *testing.T, token *Token, wantValue string, wantType TokenType) {
	assert.Equal(t, token.Typ, wantType)
	assert.Equal(t, token.Value, wantValue)
}

func TestFromFile(t *testing.T) {
	tokens := FromFile("./../tests/function.ts")
	assert.Equal(t, len(tokens), 16)
}
