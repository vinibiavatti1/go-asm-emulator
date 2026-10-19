package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Jne(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	flags := machine.EFlags().Get()
	zf, _ := FromFlags(flags)
	if !zf {
		return interpreter.JumpToLabel(label)
	}
	interpreter.IncInstructionPtr()
	return nil
}
