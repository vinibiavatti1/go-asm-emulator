package instruction

import "vinibiavatti1/goasmemulator/machine"

func Jge(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	label := instruction.Arg1
	flags := machine.EFlags().Get()
	_, sf := FromFlags(flags)
	if !sf {
		return interpreter.JumpToLabel(label)
	}
	interpreter.IncInstructionPtr()
	return nil
}
