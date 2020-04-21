package operand

import "fmt"

var _ Operand = &Offset{}

type Offset struct {
	Offset int
}

func NewOffset(offset int) *Offset {
	return &Offset{Offset: offset}
}

func (o *Offset) String() string {
	return fmt.Sprintf("%d", o.Offset)
}

func (*Offset) Typ() OperandType {
	return TYPE_OFFSET
}
