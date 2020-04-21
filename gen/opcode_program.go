package gen

import (
	"strconv"
	"strings"
	"tinyscript/translator"
)

type OpCodeProgram struct {
	Entry        *int
	Instructions []*Instruction
	Comments     map[int]string
}

func NewOpCodeProgram() *OpCodeProgram {
	return &OpCodeProgram{Entry: nil, Instructions: make([]*Instruction, 0), Comments: make(map[int]string)}
}

func (o *OpCodeProgram) Add(instr *Instruction) {
	o.Instructions = append(o.Instructions, instr)
}
func (o *OpCodeProgram) String() string {
	prts := make([]string, 0, len(o.Instructions))
	for i, instr := range o.Instructions {
		if c, ok := o.Comments[i]; ok {
			prts = append(prts, "#"+c)
		}
		str := instr.String()
		if o.Entry != nil && *o.Entry == i {
			str = "MAIN:" + str
		}
		prts = append(prts, str)
	}

	return strings.Join(prts, "\n")
}

func (o *OpCodeProgram) SetEntry(entry *int) {
	o.Entry = entry
}

func (o *OpCodeProgram) AddComment(comment string) {
	o.Comments[len(o.Comments)] = comment
}

func (o *OpCodeProgram) ToByteCode() []int {
	codes := []int{}
	for _, instr := range o.Instructions {
		codes = append(codes, instr.ToByteCode())
	}

	return codes
}

func (o *OpCodeProgram) GetStaticArea(taProgram *translator.TAProgram) []int {
	l := []int{}
	for _, symbol := range taProgram.StaticTable.Symbols {
		i, err := strconv.Atoi(symbol.Lexeme.Value)
		if nil != err {
			panic(err)
		}
		l = append(l, i)
	}
	return l
}
