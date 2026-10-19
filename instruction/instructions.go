package instruction

import (
	"fmt"
	"vinibiavatti1/goasmemulator/machine"
)

type InstructionEvaluator func(*machine.Machine, *Interpreter, *Instruction) error

func EvaluateInstruction(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	evaluator, ok := instructions[instruction.Name]
	if !ok {
		return fmt.Errorf("Instruction not found: %s", instruction.Name)
	}
	return evaluator(machine, interpreter, instruction)
}

var instructions = map[string]InstructionEvaluator{
	"mov":  Mov,
	"push": Push,
	"pop":  Pop,
	"add":  Add,
	"sub":  Sub,
	"inc":  Inc,
	"dec":  Dec,
	"not":  Not,
	"neg":  Neg,
	"and":  And,
	"or":   Or,
	"xor":  Xor,
	"shl":  Shl,
	"shr":  Shr,
	"jmp":  Jmp,
	"cmp":  Cmp,
	"je":   Je,
	"jz":   Je,
	"jne":  Jne,
	"jnz":  Jne,
	"jl":   Jl,
	"jge":  Jge,
	"jle":  Jle,
	"jg":   Jg,
	"call": Call,
	"ret":  Ret,
	"nop":  Nop,
}
