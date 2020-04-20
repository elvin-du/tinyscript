package symbol

import (
	"fmt"
	"tinyscript/lexer"
)

type Table struct {
	Parent      *Table
	Children    []*Table
	Symbols     []*Symbol
	TempIndex   int
	OffsetIndex int
	Level       int
}

func NewTable() *Table {
	return &Table{
		Symbols:  make([]*Symbol, 0),
		Children: make([]*Table, 0),
	}
}

func (t *Table) AddSymbol(symbol *Symbol) {
	t.Symbols = append(t.Symbols, symbol)
	symbol.Parent = t
}

func (t *Table) symbolByLexeme(lexeme *lexer.Token) *Symbol {
	for _, v := range t.Symbols {
		if lexeme.Value == v.Lexeme.Value {
			return v
		}
	}
	return nil
}

func (t *Table) Exists(lexeme *lexer.Token) bool {
	symbol := t.symbolByLexeme(lexeme)
	if nil != symbol {
		return true
	}

	if t.Parent != nil {
		return t.Parent.Exists(lexeme)
	}

	return false
}

func (t *Table) CloneFromSymbolTree(lexeme *lexer.Token, layoutOffset int) *Symbol {
	symbl := t.symbolByLexeme(lexeme)
	if nil != symbl {
		symbol := *symbl
		symbol.LayerOffset = layoutOffset
		return &symbol
	}
	if nil != t.Parent {
		return t.CloneFromSymbolTree(lexeme, layoutOffset+1)
	}

	return nil
}

func (t *Table) CreateSymbolByLexeme(lexeme *lexer.Token) *Symbol {
	var symbol *Symbol = nil
	if lexeme.IsScalar() {
		symbol = MakeImmediateSymbol(lexeme)
	} else {
		symbol = t.CloneFromSymbolTree(lexeme, 0)
		if symbol == nil {
			t.OffsetIndex +=1
			symbol = MakeAddressSymbol(lexeme, t.OffsetIndex)
		}
	}
	t.AddSymbol(symbol)

	return symbol
}

func (t *Table) CreateVariable() *Symbol {
	t.TempIndex += 1
	lexeme := lexer.NewToken(lexer.VARIABLE, "p"+fmt.Sprintf("%d", t.TempIndex))
	t.OffsetIndex += 1
	symbol := MakeAddressSymbol(lexeme, t.OffsetIndex)
	t.AddSymbol(symbol)
	return symbol
}

func (t *Table) AddChild(child *Table) {
	child.Parent = t
	child.Level = t.Level + 1
	t.Children = append(t.Children, child)
}

func (t *Table) LocalSize() int {
	return t.OffsetIndex
}

func (t *Table) CreateLabel(label string, lexeme *lexer.Token) {
	labelSymbol := MakeLabelSymbol(label, lexeme)
	t.AddSymbol(labelSymbol)
}
