package gen

import (
	"fmt"
	"reflect"
	"strings"
	"tinyscript/gen/operand"
	"tinyscript/translator/symbol"
)

const (
	MASK_OPCODE  = 0xfc000000
	MASK_R0      = 0x03e00000
	MASK_R1      = 0x001f0000
	MASK_R2      = 0x0000f800
	MASK_OFFSET0 = 0x03ffffff
	MASK_OFFSET1 = 0x001fffff
	MASK_OFFSET2 = 0x000007ff
)

type Instruction struct {
	Code   *OpCode
	OpList []operand.Operand
}

func NewInstruction(code *OpCode) *Instruction {
	return &Instruction{Code: code, OpList: make([]operand.Operand, 0)}
}
func NewJumpInstruction(code *OpCode, offset int) *Instruction {
	i := NewInstruction(code)
	i.AddOperand(operand.NewOffset(offset))
	return i
}

func NewOffsetInstruction(code *OpCode, r1, r2 *operand.Register, offset *operand.Offset) *Instruction {
	i := NewInstruction(code)
	i.AddOperand(r1)
	i.AddOperand(r2)
	i.AddOperand(offset)
	return i
}

func NewRegisterInstruction(code *OpCode, r1, r2, r3 *operand.Register) *Instruction {
	i := NewInstruction(code)
	i.AddOperand(r1)
	if r2 != nil {
		i.AddOperand(r2)
	}
	if r3 != nil {
		i.AddOperand(r3)
	}
	return i
}

func NewBNEInstruction(r1, r2 *operand.Register, label string) *Instruction {
	i := NewInstruction(BNE)
	i.AddOperand(r1)
	i.AddOperand(r2)
	i.AddOperand(operand.NewLabel(label))
	return i
}

func NewImmediateInstruction(code *OpCode, r1 *operand.Register, number *operand.ImmediateNumber) *Instruction {
	i := NewInstruction(code)
	i.AddOperand(r1)
	i.AddOperand(number)
	return i
}

func (i *Instruction) AddOperand(o operand.Operand) {
	i.OpList = append(i.OpList, o)
}

func (i *Instruction) String() string {
	s := i.Code.String()
	prts := make([]string, 0, len(i.OpList)+1)
	for _, op := range i.OpList {
		prts = append(prts, op.String())
	}
	return s + " " + strings.Join(prts, " ")
}

func (i *Instruction) ToByteCode() int {
	code := 0
	x := i.Code.Value
	code |= int(x) << 26
	switch i.Code.AddrType {
	case ADDRESSING_TYPE_IMMEDIATE:
		r0 := i.OpList[0].(*operand.Register)
		code |= int(r0.Addr) << 21
		code |= i.OpList[1].(*operand.ImmediateNumber).Value
		return code
	case ADDRESSING_TYPE_REGISTER:
		r1 := i.OpList[0].(*operand.Register)
		code |= int(r1.Addr) << 21
		if len(i.OpList) > 1 {
			code |= int(i.OpList[1].(*operand.Register).Addr) << 16
			if len(i.OpList) > 2 {
				r2 := int(i.OpList[2].(*operand.Register).Addr)
				code |= r2 << 11
			}
		}
	case ADDRESSING_TYPE_JUMP:
		if len(i.OpList) > 0 {
			code |= i.OpList[0].(*operand.Label).Offset.GetEncodedOffset()
		}
	case ADDRESSING_TYPE_OFFSET:
		r1 := i.OpList[0].(*operand.Register)
		r2 := i.OpList[1].(*operand.Register)
		var offset *operand.Offset = nil
		if reflect.TypeOf(i.OpList[2]).String() == reflect.TypeOf(&operand.Label{}).String() {
			offset = i.OpList[2].(*operand.Label).Offset
		} else {
			offset = i.OpList[2].(*operand.Offset)
		}

		code |= int(r1.Addr) << 21
		code |= int(r2.Addr) << 16
		code |= offset.GetEncodedOffset()
	}

	return code
}

func LoadToRegister(target *operand.Register, arg *symbol.Symbol) *Instruction {
	//转成证书，目前只支持整数
	if arg.Typ == symbol.SYMBOL_ADDRESS {
		return NewOffsetInstruction(LW, target, operand.SP, operand.NewOffset(-arg.Offset))
	} else if arg.Typ == symbol.SYMBOL_IMMEDIATE {
		return NewOffsetInstruction(LW, target, operand.STATIC, operand.NewOffset(arg.Offset))
	}

	panic(fmt.Sprintf("Cannot load type %v symbol to register", arg.Typ))
}

func SaveToMemory(source *operand.Register, arg *symbol.Symbol) *Instruction {
	return NewOffsetInstruction(SW, source, operand.SP, operand.NewOffset(-arg.Offset))
}

func FromByCode(code int) *Instruction {
	byteOpcode := (byte)(int(code&MASK_OPCODE) >> 26)
	opcode := FromByte(byteOpcode)
	i := NewInstruction(opcode)

	switch opcode.AddrType {
	case ADDRESSING_TYPE_IMMEDIATE:
		reg := (code & MASK_R0) >> 21
		number := code & MASK_OFFSET1
		i.OpList = append(i.OpList, operand.RegisterFromAddr(reg))
		i.OpList = append(i.OpList, operand.NewImmediateNumber(number))
	case ADDRESSING_TYPE_REGISTER:
		r1Addr := (code & MASK_R0) >> 21
		r2Addr := (code & MASK_R1) >> 16
		r3Addr := (code & MASK_R2) >> 11
		r1 := operand.RegisterFromAddr(r1Addr)

		var r2 *operand.Register = nil
		if r2Addr != 0 {
			r2 = operand.RegisterFromAddr(r2Addr)
		}

		var r3 *operand.Register = nil
		if r3Addr != 0 {
			r3 = operand.RegisterFromAddr(r3Addr)
		}

		i.OpList = append(i.OpList, r1)

		if nil != r2 {
			i.OpList = append(i.OpList, r2)
		}
		if nil != r3 {
			i.OpList = append(i.OpList, r3)
		}
	case ADDRESSING_TYPE_JUMP:
		offset := code & MASK_OFFSET0
		i.OpList = append(i.OpList, operand.DecodeOffset(offset))
	case ADDRESSING_TYPE_OFFSET:
		r1Addr := (code & MASK_R0) >> 21
		r2Addr := (code & MASK_R1) >> 16
		offset := code & MASK_OFFSET2
		i.OpList = append(i.OpList, operand.RegisterFromAddr(r1Addr))
		i.OpList = append(i.OpList, operand.RegisterFromAddr(r2Addr))
		i.OpList = append(i.OpList, operand.DecodeOffset(offset))
	}

	return i
}

func (i *Instruction) GetOperand(index int) operand.Operand {
	return i.OpList[index]
}
