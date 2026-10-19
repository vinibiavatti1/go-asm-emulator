package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Call(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	machine.Stack().Push(instruction.LineNumber)
	interpreter.JumpToLabel(label)
	return nil
}
