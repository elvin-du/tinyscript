package translator

import (
	"fmt"
	"tinyscript/translator/symbol"
)

type TAInstruction struct {
	Arg1   interface{}
	Arg2   interface{}
	Op     string
	Result *symbol.Symbol
	Typ    TAInstructionType
	Label  string
}

func NewTAInstruction(typ TAInstructionType, result *symbol.Symbol, op string, arg1 interface{}, arg2 interface{}) *TAInstruction {
	return &TAInstruction{Arg1: arg1, Arg2: arg2, Op: op, Result: result, Typ: typ}
}

func (t TAInstruction) String() string {
	switch t.Typ {
	case TAINSTR_TYPE_ASSIGN:
		if nil != t.Arg2 {
			return fmt.Sprintf("%v = %v %v %v", t.Result, t.Arg1, t.Op, t.Arg2)
		} else {
			return fmt.Sprintf("%v = %v", t.Result, t.Arg1)
		}
	case TAINSTR_TYPE_IF:
		return fmt.Sprintf("IF %v ELSE %v", t.Arg1, t.Arg2)
	case TAINSTR_TYPE_GOTO:
		return fmt.Sprintf("GOTO %v", t.Arg1)
	case TAINSTR_TYPE_LABEL:
		return fmt.Sprintf("%v:", t.Arg1)
	case TAINSTR_TYPE_FUNC_BEGIN:
		return "FUNC_BEGIN"
	case TAINSTR_TYPE_RETURN:
		return fmt.Sprintf("RETURN %v", t.Arg1)
	case TAINSTR_TYPE_PARAM:
		return fmt.Sprintf("PARAM %v %v", t.Arg1, t.Arg2)
	case TAINSTR_TYPE_SP:
		return fmt.Sprintf("SP %v", t.Arg1)
	case TAINSTR_TYPE_CALL:
		return fmt.Sprintf("CALL %v", t.Arg1)
	}

	panic("unknown opcode type")
}
