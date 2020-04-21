package operand

import "fmt"

var _ Operand = &Register{}

var (
	Registers = [31]*Register{}

	ZERO   = NewRegister("ZERO", 1)
	PC     = NewRegister("PC", 2)
	SP     = NewRegister("SP", 3)
	STATIC = NewRegister("STATIC", 4)
	RA     = NewRegister("RA", 5)

	S0 = NewRegister("S0", 10)
	S1 = NewRegister("S1", 11)
	S2 = NewRegister("S2", 12)

	L0 = NewRegister("L0", 20)
)

type Register struct {
	Addr byte
	Name string
}

func NewRegister(name string, addr byte) *Register {
	reg := &Register{Addr: addr, Name: name}
	Registers[addr] = reg
	return reg
}

func (reg *Register) Typ() OperandType {
	return TYPE_REGISTER
}
func (reg *Register) String() string {
	return reg.Name
}

func RegisterFromAddr(reg int) *Register {
	if reg < 0 || reg >= len(Registers) {
		panic(fmt.Sprintf("no register's address is %d", reg))
	}

	return Registers[reg]
}
