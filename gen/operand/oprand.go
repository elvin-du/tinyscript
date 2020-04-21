package operand

type Operand interface {
	String() string
	Typ() OperandType
}
