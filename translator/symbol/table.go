package symbol

import "tinyscript/lexer"

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
	symbl := t.symbolByLexeme(lexeme)
	if nil != symbl {
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
			symbol = MakeAddressSymbol(lexeme, t.OffsetIndex+1)
		}
	}

	t.Symbols = append(t.Symbols, symbol)

	return symbol
}

