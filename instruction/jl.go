package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Jl(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	flags := machine.EFlags().Get()
	_, sf := FromFlags(flags)
	if sf {
		return interpreter.JumpToLabel(label)
	}
	interpreter.IncInstructionPtr()
	return nil
}
