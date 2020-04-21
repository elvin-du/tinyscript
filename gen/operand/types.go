package operand

type OperandType int

const (
	TYPE_REGISTER = iota
	TYPE_IMMEDIATE
	TYPE_LABEL
	TYPE_OFFSET
)
