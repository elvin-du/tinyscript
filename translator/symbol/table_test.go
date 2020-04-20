package symbol

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/lexer"
)

func TestSymbolTable(t *testing.T) {
	table := NewTable()
	table.CreateLabel("L0", lexer.NewToken(lexer.VARIABLE, "foo"))
	table.CreateVariable()
	table.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "foo"))
	assert.Equal(t, table.LocalSize(), 1)
}

func TestTableChain(t *testing.T) {
	table := NewTable()
	table.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "a"))

	childTable := NewTable()
	table.AddChild(childTable)

	childChildTable := NewTable()
	childTable.AddChild(childChildTable)

	assert.Equal(t, childChildTable.Exists(lexer.NewToken(lexer.VARIABLE, "a")), true)
	assert.Equal(t, childTable.Exists(lexer.NewToken(lexer.VARIABLE, " a")), true)
}

func TestOffset(t *testing.T) {
	table := NewTable()

	table.CreateSymbolByLexeme(lexer.NewToken(lexer.INTEGER, "100"))
	symbola := table.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "a"))
	symbolb := table.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "b"))

	childTable := NewTable()
	table.AddChild(childTable)
	anotherSymbolB := childTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "b"))
	symbolC := childTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "c"))

	assert.Equal(t, symbola.Offset, 0)
	assert.Equal(t, symbolb.Offset, 1)
	assert.Equal(t, anotherSymbolB.Offset, 1)
	assert.Equal(t, anotherSymbolB.LayerOffset, 1)
	assert.Equal(t, symbolC.Offset, 0)
	assert.Equal(t, symbolC.LayerOffset, 0)
}
