package operand

import "fmt"

var _ Operand = &ImmediateNumber{}

type ImmediateNumber struct {
	Value int
}

func NewImmediateNumber(value int) *ImmediateNumber {
	return &ImmediateNumber{Value: value}
}

func (i *ImmediateNumber) String() string {
	return fmt.Sprintf("%d", i.Value)
}

func (*ImmediateNumber) Typ() OperandType {
	return TYPE_IMMEDIATE
}
