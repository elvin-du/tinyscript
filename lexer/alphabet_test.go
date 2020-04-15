package lexer

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAlphabet(t *testing.T) {
	assert.Equal(t, IsLetter("a"), true)
	assert.Equal(t, IsLiteral("a"), true)
	assert.Equal(t, IsNumber("2"), true)
	assert.Equal(t, IsOperator("*"), true)
	assert.Equal(t, IsOperator("^"), true)
	assert.Equal(t, IsOperator("-"), true)
	assert.Equal(t, IsOperator("="), true)
	assert.Equal(t, IsOperator("/"), true)
	assert.Equal(t, IsOperator("%"), true)
}
