package translator

import (
	"fmt"
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

func (t *TAProgram) AddLabel() *TAInstruction {
	label := fmt.Sprintf("L%d", t.LabelCounter)
	t.LabelCounter += 1
	taCode := NewTAInstruction(TAINSTR_TYPE_LABEL, nil, "", nil, nil)
	taCode.Arg1 = label
	t.Instructions = append(t.Instructions, taCode)
	return taCode
}

func (t *TAProgram) String() string {
	var lines []string
	for _, v := range t.Instructions {
		lines = append(lines, v.String())
	}

	return strings.Join(lines, "\n")
}

//根据符号表的内容，判断符号类型，如果是SYMBOL_IMMEDIATE，则加入静态符号表，以此来设置静态符号表的信息
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
