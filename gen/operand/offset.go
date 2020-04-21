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

func (o *Offset) GetEncodedOffset() int {
	if o.Offset > 0 {
		return o.Offset
	}

	return 0x400 | -o.Offset
}
func DecodeOffset(offset int) *Offset {
	if offset&0x400 > 0 {
		offset = offset & 0x3ff
		offset = -offset
	}
	return NewOffset(offset)
}
func (*Offset) Typ() OperandType {
	return TYPE_OFFSET
}
