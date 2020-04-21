package gen

import (
	"strings"
	"tinyscript/gen/operand"
)

type Instruction struct {
	Code   *OpCode
	OpList []operand.Operand
}

func (i *Instruction) String() string {
	s := i.Code.String()
	prts := make([]string, 0, len(i.OpList)+1)
	for _, op := range i.OpList {
		prts = append(prts, op.String())
	}
	return s + " " + strings.Join(prts, " ")
}
