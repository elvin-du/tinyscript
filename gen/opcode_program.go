package gen

import "strings"

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
