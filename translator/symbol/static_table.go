package symbol

import (
	"fmt"
	"strings"
)

type StaticSymbolTable struct {
	OffsetMap     map[string]*Symbol
	OffsetCounter int
	Symbols       []*Symbol
}

func NewStaticSymbolTable() *StaticSymbolTable {
	return &StaticSymbolTable{OffsetCounter: 0, OffsetMap: make(map[string]*Symbol), Symbols: make([]*Symbol, 0)}
}

func (s *StaticSymbolTable) Add(symbol *Symbol) {
	lexval := symbol.Lexeme.Value
	if _, ok := s.OffsetMap[lexval]; !ok {
		s.OffsetMap[lexval] = symbol
		symbol.Offset = s.OffsetCounter
		s.OffsetCounter += 1
		s.Symbols = append(s.Symbols, symbol)
	} else {
		sameSymbol := s.OffsetMap[lexval]
		symbol.Offset = sameSymbol.Offset
	}
}

func (s *StaticSymbolTable) Size() int {
	return len(s.Symbols)
}

func (s *StaticSymbolTable) String() string {
	var list []string
	for i, v := range s.Symbols {
		list = append(list, fmt.Sprintf("%d:%s", i, v))
	}

	return strings.Join(list, "\n")
}
