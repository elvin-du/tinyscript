package gen

import (
	"fmt"
	"tinyscript/gen/operand"
	"tinyscript/translator"
	"tinyscript/translator/symbol"
)

type OpCodeGen struct {
}

func NewOpCodeGen() *OpCodeGen {
	return &OpCodeGen{}
}

func (g *OpCodeGen) Gen(taProgram *translator.TAProgram) *OpCodeProgram {
	program := NewOpCodeProgram()
	taInstrs := taProgram.Instructions
	labelHash := make(map[string]int)

	for _, taInstr := range taInstrs {
		program.AddComment(taInstr.String())
		switch taInstr.Typ {
		case translator.TAINSTR_TYPE_ASSIGN:
			g.GenCopy(program, taInstr)
		case translator.TAINSTR_TYPE_GOTO:
			g.GenGoTo(program, taInstr)
		case translator.TAINSTR_TYPE_CALL:
			g.GenCopy(program, taInstr)
		case translator.TAINSTR_TYPE_PARAM:
			g.GenPass(program, taInstr)
		case translator.TAINSTR_TYPE_SP:
			g.GenSP(program, taInstr)
		case translator.TAINSTR_TYPE_LABEL:
			if taInstr.Arg2 != nil && taInstr.Arg2.(string) == "main" {
				size := len(program.Instructions)
				program.SetEntry(&size)
				labelHash[taInstr.Arg2.(string)] = len(program.Instructions)
			}
		case translator.TAINSTR_TYPE_RETURN:
			g.GenReturn(program, taInstr)
		case translator.TAINSTR_TYPE_FUNC_BEGIN:
			g.GenFuncBegin(program, taInstr)
		case translator.TAINSTR_TYPE_IF:
			g.GenIf(program, taInstr)
		default:
			panic(fmt.Sprintf("unknown type %d", taInstr.Typ))
		}
	}

	g.Relabel(program, labelHash)

	return program
}

func (g *OpCodeGen) GenGoTo(program *OpCodeProgram, ta *translator.TAInstruction) {
	label := ta.Arg1.(string)
	i := NewInstruction(JUMP)
	//label对应的未知在relabel阶段计算
	i.OpList = append(i.OpList, operand.NewLabel(label))
	program.Add(i)
}

func (g *OpCodeGen) GenIf(program *OpCodeProgram, ta *translator.TAInstruction) {
	label := ta.Arg2
	program.Add(NewBNEInstruction(operand.S2, operand.ZERO, label.(string)))
}
func (g *OpCodeGen) Relabel(program *OpCodeProgram, labelMap map[string]int) {
	for _, instr := range program.Instructions {
		if instr.Code == JUMP || instr.Code == JR || instr.Code == BNE {
			idx := 0
			if instr.Code == BNE {
				idx = 2
			}
			labelOperand := instr.OpList[idx].(*operand.Label)
			label := labelOperand.Label
			offset := labelMap[label]
			labelOperand.Offset.Offset = offset
		}
	}
}

func (g *OpCodeGen) GenReturn(program *OpCodeProgram, ta *translator.TAInstruction) {
	if nil != ta.Arg1 {
		ret := ta.Arg1.(*symbol.Symbol)
		program.Add(LoadToRegister(operand.S0, ret))
	}
	program.Add(NewOffsetInstruction(SW, operand.S0, operand.SP, operand.NewOffset(1)))
	i := NewInstruction(RETURN)
	program.Add(i)
}

func (g *OpCodeGen) GenSP(program *OpCodeProgram, ta *translator.TAInstruction) {
	offset := ta.Arg1.(int)
	if offset > 0 {
		program.Add(NewImmediateInstruction(ADDI, operand.SP, operand.NewImmediateNumber(offset)))
	} else {
		program.Add(NewImmediateInstruction(SUBI, operand.SP, operand.NewImmediateNumber(-offset)))
	}
}

func (g *OpCodeGen) GenPass(program *OpCodeProgram, ta *translator.TAInstruction) {
	arg1 := ta.Arg1.(*symbol.Symbol)
	no := ta.Arg2.(int)
	program.Add(LoadToRegister(operand.S0, arg1))
	//pass a
	program.Add(NewOffsetInstruction(SW, operand.S0, operand.SP, operand.NewOffset(-no)))
}

func (g *OpCodeGen) GenFuncBegin(program *OpCodeProgram, ta *translator.TAInstruction) {
	i := NewOffsetInstruction(SW, operand.RA, operand.SP, operand.NewOffset(0))
	program.Add(i)
}

func (g *OpCodeGen) GenCall(program *OpCodeProgram, ta *translator.TAInstruction) {
	label := ta.Arg1.(*symbol.Symbol)
	i := NewInstruction(JR)
	i.OpList = append(i.OpList, operand.NewLabel(label.Label))
	program.Add(i)
}

func (g *OpCodeGen) GenCopy(program *OpCodeProgram, ta *translator.TAInstruction) {
	result := ta.Result
	op := ta.Op
	arg1 := ta.Arg1.(*symbol.Symbol)

	if nil == ta.Arg2 {
		program.Add(LoadToRegister(operand.S0, arg1))
		program.Add(SaveToMemory(operand.S0, result))
	} else {
		program.Add(LoadToRegister(operand.S0, arg1))
		arg2 := ta.Arg2.(*symbol.Symbol)
		program.Add(LoadToRegister(operand.S1, arg2))
		switch op {
		case "+":
			program.Add(NewRegisterInstruction(ADD, operand.S2, operand.S0, operand.S1))
		case "-":
			program.Add(NewRegisterInstruction(SUB, operand.S2, operand.S0, operand.S1))
		case "*":
			program.Add(NewRegisterInstruction(MULT, operand.S0, operand.S1, nil))
			program.Add(NewRegisterInstruction(MFLO, operand.S2, nil, nil))
		case "==":
			program.Add(NewRegisterInstruction(EQ, operand.S2, operand.S1, operand.S0))
		}
		program.Add(SaveToMemory(operand.S2, result))
	}
}
