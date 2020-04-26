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
	taProg := translator.NewTranslator().Translate(parser.ParseFromFile("../tests/fact2.ts"))
	t.Log(taProg)
	prog := gen.NewOpCodeGen().Gen(taProg)
	staticTable := prog.GetStaticArea(taProg)
	opcodes := prog.ToByteCode()
	t.Log(prog)
	t.Log(taProg.StaticTable)
	vm := NewVM(staticTable, opcodes, prog.Entry)
	// CALL main
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	t.Log("RA:", vm.Registers[operand.RA.Addr])
	assert.Equal(t, vm.GetSpMemory(0), 39)

	// PARAM 10 0
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, vm.GetSpMemory(-3), 2)

	// SP -2
	vm.runOneStep()
	vm.runOneStep()
	t.Log("RA:", vm.Registers[operand.RA.Addr])

	// #FUNC_BEGIN
	vm.runOneStep()
	assert.Equal(t, vm.GetSpMemory(0), 33)

	// #p1 = n == 0
	assert.Equal(t, vm.GetSpMemory(-1), 2)
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, vm.GetSpMemory(-2) == 0, false)

	// #IF p1 ELSE L1
	vm.runOneStep()
	vm.runOneStep()

	// #p3 = n - 1
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, 1, vm.GetSpMemory(-3))

	// #PARAM p3 0
	// #SP-5
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, 1, vm.GetSpMemory(-1))

	vm.runOneStep()
	vm.runOneStep()

	// #p1 = n == 0
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, false, vm.GetSpMemory(-2) == 0)

	// #IF p1 ELSE L1
	vm.runOneStep()

	// #p3 = n - 1
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()

	// #PARAM p3 0
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()

	// CALL
	vm.runOneStep()
	vm.runOneStep()

	// #p1 = n == 0
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	assert.Equal(t, true, vm.GetSpMemory(-2) == 0)

	// #IF p1 ELSE L1
	vm.runOneStep()

	// RETURN 1
	vm.runOneStep()
	vm.runOneStep()

	vm.runOneStep()
	vm.runOneStep()

	// #p4 = p2 * n 计算递归值
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	// #RETURN p4
	vm.runOneStep()
	vm.runOneStep()
	//RETURN
	vm.runOneStep()
	vm.runOneStep()

	//#p4 = p2 * n
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()
	vm.runOneStep()

	assert.Equal(t, 2, vm.GetSpMemory(-5))

	vm.runOneStep()
	vm.runOneStep()
	// RETURN MAIN
	vm.runOneStep()

	// SP 2
	vm.runOneStep()

	// #RETURN p1 : from main
	vm.runOneStep()
	assert.Equal(t, 2, vm.GetSpMemory(-1))

	for ; vm.runOneStep(); {
	}
	assert.Equal(t, 2, vm.GetSpMemory(0))
}

func TestRecursivefunction1(t *testing.T) {
	taProg := translator.NewTranslator().Translate(parser.ParseFromFile("../tests/fact5.ts"))
	prog := gen.NewOpCodeGen().Gen(taProg)
	staticTable := prog.GetStaticArea(taProg)
	opcodes := prog.ToByteCode()
	//t.Log(prog)
	//t.Log(taProg.StaticTable)
	vm := NewVM(staticTable, opcodes, prog.Entry)
	vm.run()
	assert.Equal(t, 120, vm.GetSpMemory(0))
}
