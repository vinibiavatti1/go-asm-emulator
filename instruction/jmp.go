package instruction

import "vinibiavatti1/goasmemulator/machine"

func Jmp(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	interpreter.JumpToLabel(label)
	return nil
}
