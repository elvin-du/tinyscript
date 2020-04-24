package gen

import "fmt"

var Codes = [63]*OpCode{}

var (
	ADD  = NewOpCode(ADDRESSING_TYPE_REGISTER, "ADD", 0x01)
	SUB  = NewOpCode(ADDRESSING_TYPE_REGISTER, "SUB", 0x02)
	MULT = NewOpCode(ADDRESSING_TYPE_REGISTER, "MULT", 0x03)

	ADDI  = NewOpCode(ADDRESSING_TYPE_IMMEDIATE, "ADDI", 0x05) //立即数加
	SUBI  = NewOpCode(ADDRESSING_TYPE_IMMEDIATE, "SUBI", 0x06)
	MULTI = NewOpCode(ADDRESSING_TYPE_IMMEDIATE, "MULTI", 0x07)

	MFLO = NewOpCode(ADDRESSING_TYPE_REGISTER, "MFLO", 0x08) //MULT/MULTI操作码的结果会存储到这个寄存器中

	EQ  = NewOpCode(ADDRESSING_TYPE_REGISTER, "EQ", 0x09)
	BNE = NewOpCode(ADDRESSING_TYPE_OFFSET, "BNE", 0x15) //不相等

	SW = NewOpCode(ADDRESSING_TYPE_OFFSET, "SW", 0x10) //从寄存器写回内存
	LW = NewOpCode(ADDRESSING_TYPE_OFFSET, "LW", 0x11) //从内存读入到寄存器

	JUMP   = NewOpCode(ADDRESSING_TYPE_JUMP, "JUMP", 0x20)
	JR     = NewOpCode(ADDRESSING_TYPE_JUMP, "JR", 0x21) //函数的跳转
	RETURN = NewOpCode(ADDRESSING_TYPE_JUMP, "RETURN", 0x22)
)

type OpCode struct {
	Name     string
	Value    byte
	AddrType AddressingType
}

func NewOpCode(addrType AddressingType, name string, value byte) *OpCode {
	oc := &OpCode{Name: name, Value: value, AddrType: addrType}
	Codes[value] = oc
	return oc
}

func (oc *OpCode) String() string {
	return oc.Name
}

func FromByte(byteOpcode byte) *OpCode {
	code := Codes[byteOpcode]
	if nil == code {
		panic(fmt.Sprintf("%x opcode undefined", byteOpcode))
	}

	return code
}
