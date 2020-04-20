package translator

import (
	"strings"
	"tinyscript/translator/symbol"
)

type TAProgram struct {
	Instructions []*TAInstruction
	LabelCounter int
	StaticTable  *symbol.StaticSymbolTable
}

func NewTAProgram() *TAProgram {
	return &TAProgram{Instructions: make([]*TAInstruction, 0), StaticTable: symbol.NewStaticSymbolTable()}
}

func (t *TAProgram) Add(instr *TAInstruction) {
	t.Instructions = append(t.Instructions, instr)
}

func (t *TAProgram) String() string {
	var lines []string
	for _, v := range t.Instructions {
		lines = append(lines, v.String())
	}

	return strings.Join(lines, "\n")
}

func (t *TAProgram) SetStaticSymbols(table *symbol.Table) {
	for _, v := range table.Symbols {
		if symbol.SYMBOL_IMMEDIATE == v.Typ {
			t.StaticTable.Add(v)
		}
	}

	for _, child := range table.Children {
		t.SetStaticSymbols(child)
	}
}
