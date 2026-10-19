package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Jle(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	flags := machine.EFlags().Get()
	zf, sf := FromFlags(flags)
	if zf || sf {
		return interpreter.JumpToLabel(label)
	}
	interpreter.IncInstructionPtr()
	return nil
}
