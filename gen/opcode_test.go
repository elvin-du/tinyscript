package gen

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
	"tinyscript/gen/operand"
	symbol2 "tinyscript/translator/symbol"
)

func TestAdd(t *testing.T) {
	a := NewInstruction(ADD)
	a.AddOperand(operand.S2)
	a.AddOperand(operand.S0)
	a.AddOperand(operand.S1)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}
func TestMult(t *testing.T) {
	a := NewInstruction(MULT)
	a.AddOperand(operand.S0)
	a.AddOperand(operand.S1)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestNewJumpInstruction(t *testing.T) {
	a := NewInstruction(JUMP)
	label := operand.NewLabel("L0")
	a.AddOperand(label)
	label.SetOffset(100)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestJR(t *testing.T) {
	a := NewInstruction(JR)
	label := operand.NewLabel("L0")
	a.AddOperand(label)
	label.SetOffset(100)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestSW(t *testing.T) {
	symbol := symbol2.NewSymbol(symbol2.SYMBOL_IMMEDIATE)
	symbol.Offset = -100
	a := SaveToMemory(operand.S0, symbol)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestSW1(t *testing.T) {
	symbol := symbol2.NewSymbol(symbol2.SYMBOL_IMMEDIATE)
	symbol.Offset = 100
	a := SaveToMemory(operand.S0, symbol)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestLW(t *testing.T) {
	symbol := symbol2.NewSymbol(symbol2.SYMBOL_IMMEDIATE)
	symbol.Offset = 100
	a := LoadToRegister(operand.S0, symbol)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestLW2(t *testing.T) {
	symbol := symbol2.NewSymbol(symbol2.SYMBOL_ADDRESS)
	symbol.Offset = 100
	a := LoadToRegister(operand.S0, symbol)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestSP(t *testing.T) {
	a := NewImmediateInstruction(ADDI, operand.SP, operand.NewImmediateNumber(100))
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func TestBNE(t *testing.T) {
	a := NewBNEInstruction(operand.S0, operand.S1, "L0")
	a.GetOperand(2).(*operand.Label).SetOffset(100)
	AssertSameInstruction(t, a, FromByCode(a.ToByteCode()))
}

func AssertSameInstruction(t *testing.T, a, b *Instruction) {
	assert.Equal(t, a.Code, b.Code)
	assert.Equal(t, len(a.OpList), len(b.OpList))
	for i, av := range a.OpList {
		bv := b.GetOperand(i)
		assert.Equal(t, bv, av)

		if reflect.ValueOf(av).Type().String() == reflect.TypeOf(&operand.ImmediateNumber{}).String() {
			assert.Equal(t, av.(*operand.ImmediateNumber).Value, bv.(*operand.ImmediateNumber).Value)
		} else if reflect.ValueOf(av).Type().String() == reflect.TypeOf(&operand.Offset{}).String() {
			assert.Equal(t, av.(*operand.Offset).Offset, bv.(*operand.Offset).Offset)
		} else if reflect.ValueOf(av).Type().String() == reflect.TypeOf(&operand.Register{}).String() {
			assert.Equal(t, av.(*operand.Register).Addr, bv.(*operand.Register).Addr)
			assert.Equal(t, av.(*operand.Register).Name, bv.(*operand.Register).Name)
		} else if reflect.ValueOf(av).Type().String() == reflect.TypeOf(&operand.Label{}).String() {
			assert.Equal(t, av.(*operand.Label).Offset, bv.(*operand.Label).Offset)
		} else {
			panic("unsupported encode/decode type" + av.String())
		}
	}
}
