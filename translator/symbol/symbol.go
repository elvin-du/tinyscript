package symbol

import "tinyscript/lexer"

type Symbol struct {
	Parent      *Table
	Lexeme      *lexer.Token
	Label       string
	Offset      int
	LayerOffset int
	Typ         SymbolType
}

func NewSymbol(typ SymbolType) *Symbol {
	return &Symbol{Typ: typ}
}

func MakeAddressSymbol(lexeme *lexer.Token, offset int) *Symbol {
	syb := NewSymbol(SYMBOL_ADDRESS)
	syb.Lexeme = lexeme
	syb.Offset = offset

	return syb
}

func MakeImmediateSymbol(lexeme *lexer.Token) *Symbol {
	syb := NewSymbol(SYMBOL_IMMEDIATE)
	syb.Lexeme = lexeme

	return syb
}

func MakeLabelSymbol(lexeme *lexer.Token, label string) *Symbol {
	syb := NewSymbol(SYMBOL_LABEL)
	syb.Lexeme = lexeme
	syb.Label = label

	return syb
}
