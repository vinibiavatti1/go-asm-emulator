package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Jmp(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	interpreter.JumpToLabel(label)
	return nil
}
