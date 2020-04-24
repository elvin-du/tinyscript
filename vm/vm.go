package vm

import (
	"log"
	"tinyscript/gen"
	"tinyscript/gen/operand"
)

type VM struct {
	Registers         [31]int
	Memory            [4096]int
	EndProgramSection int
	StartProgram      int
}

func NewVM(staticArea []int, opcodes []int, entry *int) *VM {
	vm := &VM{}

	i := 0
	for ; i < len(staticArea); i++ {
		vm.Memory[i] = staticArea[i]
	}

	j := i
	vm.StartProgram = i
	//mainStart := *entry + i
	for ; i < len(opcodes)+j; i++ {
		vm.Memory[i] = opcodes[i-j]
	}

	vm.Registers[operand.PC.Addr] = i - 3
	vm.EndProgramSection = i

	vm.Registers[operand.SP.Addr] = 4095
	return vm
}
func (vm *VM) Fetch() int {
	pc := vm.Registers[operand.PC.Addr]
	return vm.Memory[pc]
}

func (vm *VM) Decode(code int) *gen.Instruction {
	return gen.FromByCode(code)
}

func (vm *VM) Exec(instr *gen.Instruction) {
	code := instr.Code.Value
	log.Println("exec:", instr)

	switch code {
	case 0x01: //ADD
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		r2 := instr.GetOperand(2).(*operand.Register)
		vm.Registers[r0.Addr] = vm.Registers[r1.Addr] + vm.Registers[r2.Addr]
	case 0x09: //
	case 0x02: //SUB
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		r2 := instr.GetOperand(2).(*operand.Register)
		vm.Registers[r0.Addr] = vm.Registers[r1.Addr] - vm.Registers[r2.Addr]
	case 0x03: //MULT
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		vm.Registers[operand.L0.Addr] = vm.Registers[r0.Addr] * vm.Registers[r1.Addr]
	case 0x05: //ADDI
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.ImmediateNumber)
		vm.Registers[r0.Addr] += r1.Value
	case 0x06: //SUBI
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.ImmediateNumber)
		vm.Registers[r0.Addr] -= r1.Value
	//case 0x07: //MULI
	case 0x08: //MFLO
		r0 := instr.GetOperand(0).(*operand.Register)
		vm.Registers[r0.Addr] = vm.Registers[operand.L0.Addr]
	case 0x10: //SW
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		offset := instr.GetOperand(2).(*operand.Offset)
		R1VAL := vm.Registers[r1.Addr]
		vm.Memory[R1VAL+offset.Offset] = vm.Registers[r0.Addr]
	case 0x11: //LW
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		offset := instr.GetOperand(2).(*operand.Offset)
		R1VAL := vm.Registers[r1.Addr]
		vm.Registers[r0.Addr] = vm.Memory[R1VAL+offset.Offset]
	case 0x15: //BNE
		r0 := instr.GetOperand(0).(*operand.Register)
		r1 := instr.GetOperand(1).(*operand.Register)
		offset := instr.GetOperand(2).(*operand.Offset)
		if vm.Registers[r0.Addr] != vm.Registers[r1.Addr] {
			vm.Registers[operand.PC.Addr] = offset.Offset + vm.StartProgram - 1
		}
	case 0x20: //JUMP
		r0 := instr.GetOperand(0).(*operand.Offset)
		vm.Registers[operand.PC.Addr] = r0.Offset + vm.StartProgram - 1
	case 0x21: //JR
		r0 := instr.GetOperand(0).(*operand.Offset)
		vm.Registers[operand.RA.Addr] = vm.Registers[operand.PC.Addr]
		vm.Registers[operand.PC.Addr] = r0.Offset + vm.StartProgram - 1
	case 0x22: //RETRUN
		if instr.GetOperand(0) != nil {
			//match 返回值
		}

		spVal := vm.Registers[operand.SP.Addr]
		vm.Registers[operand.PC.Addr] = vm.Memory[spVal]
	}
}

func (vm *VM) run() {
	//模拟CPU循环
	// fetch
	// decode
	// exec
	// pc++
	for ;vm.runOneStep();{}
}

func (vm *VM) GetSpMemory(offset int) int {
	sp := vm.Registers[operand.SP.Addr]
	return vm.Memory[sp+offset]
}

func (vm *VM) runOneStep() bool {
	code := vm.Fetch()
	instr := vm.Decode(code)
	vm.Exec(instr)
	vm.Registers[operand.PC.Addr] += 1
	return vm.Registers[operand.PC.Addr] < vm.EndProgramSection
}
