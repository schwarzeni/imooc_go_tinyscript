package translator

import (
	"fmt"
	"go-tinyscript/translator/symbol"
)

// TAInstruction 三地址指令结构体
type TAInstruction struct {
	arg1   interface{}
	arg2   interface{}
	op     string
	result *symbol.Symbol
	t      TAInstructionType
	label  string
}

func (i *TAInstruction) Result() *symbol.Symbol {
	return i.result
}

func (i *TAInstruction) SetResult(result *symbol.Symbol) {
	i.result = result
}

func (i *TAInstruction) T() TAInstructionType {
	return i.t
}

func (i *TAInstruction) Op() string {
	return i.op
}

func (i *TAInstruction) Arg2() interface{} {
	return i.arg2
}

func (i *TAInstruction) SetArg2(arg2 interface{}) {
	i.arg2 = arg2
}

func (i *TAInstruction) Arg1() interface{} {
	return i.arg1
}

func (i *TAInstruction) SetArg1(arg1 interface{}) {
	i.arg1 = arg1
}

func NewTAInstruction(arg1 interface{}, arg2 interface{}, op string, result *symbol.Symbol, t TAInstructionType) *TAInstruction {
	return &TAInstruction{arg1: arg1, arg2: arg2, op: op, result: result, t: t}
}

func (i *TAInstruction) String() string {
	switch i.t {
	case ASSIGN:
		if i.arg2 != nil {
			return fmt.Sprintf("%s = %s %s %s", i.result, i.arg1, i.op, i.arg2)
		}
		return fmt.Sprintf("%s = %s", i.result, i.arg1)
	case IF:
		return fmt.Sprintf("IF %s ELSE %s", i.arg1, i.arg2)
	case GOTO:
		return fmt.Sprintf("GOTO %s", i.arg1)
	case LABEL:
		return fmt.Sprintf("%s:", i.arg1)
	case FUNC_BEGIN:
		return "FUNC_BEGIN"
	case RETURN:
		return fmt.Sprintf("RETURN %s", i.arg1)
	case PARAM:
		return fmt.Sprintf("PARAM %s %d", i.arg1, i.arg2)
	case SP:
		return fmt.Sprintf("SP %d", i.arg1)
	case CALL:
		return fmt.Sprintf("CALL %s", i.arg1)
	}
	panic("unknown opcode type: " + string(i.t))
}
