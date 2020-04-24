package vm

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"tinyscript/gen"
	"tinyscript/gen/operand"
	"tinyscript/parser"
	"tinyscript/translator"
)

func TestCalcExpr(t *testing.T) {
	source := `func main()int{var a = 2 * 3 + 4 
return
}`
	taProg := translator.NewTranslator().Translate(parser.Parse(source))
	prog := gen.NewOpCodeGen().Gen(taProg)
	staticTable := prog.GetStaticArea(taProg)
	opcodes := prog.ToByteCode()
	vm := NewVM(staticTable, opcodes, prog.Entry)

	// CALL main
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()

	t.Log("RA:", vm.Registers[operand.RA.Addr])
	assert.Equal(t, vm.GetSpMemory(0), 18)

	// p0 = 2 * 3
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, vm.Registers[operand.S0.Addr], 2)
	assert.Equal(t, vm.Registers[operand.S1.Addr], 3)
	assert.Equal(t, vm.Registers[operand.L0.Addr], 6)
	assert.Equal(t, vm.Registers[operand.S2.Addr], 6)
	assert.Equal(t, vm.GetSpMemory(-2), 6)

	// p1 = p0 + 4
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, vm.Registers[operand.S0.Addr], 6)
	assert.Equal(t, vm.Registers[operand.S1.Addr], 4)
	assert.Equal(t, vm.Registers[operand.S2.Addr], 10)
	assert.Equal(t, vm.GetSpMemory(-3), 10)

	// a = p1
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, vm.GetSpMemory(-1), 10)
	assert.Equal(t, vm.Registers[operand.S0.Addr], 10)

	// RETURN null
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()

	t.Log("SP:", vm.Registers[operand.SP.Addr])
}

func TestRecursiveFunction(t *testing.T) {
	taProg := translator.NewTranslator().Translate(parser.Parse("../tests/fact2.ts"))
	prog := gen.NewOpCodeGen().Gen(taProg)
	staticTable := prog.GetStaticArea(taProg)
	opcodes := prog.ToByteCode()
	t.Log(taProg.StaticTable)
	vm := NewVM(staticTable, opcodes, prog.Entry)
	// CALL main
	vm.runOneStep();
	vm.runOneStep();
	vm.runOneStep();
	t.Log("RA:", vm.Registers[operand.RA.Addr])
	assert.Equal(t, vm.GetSpMemory(0), 18)
}
