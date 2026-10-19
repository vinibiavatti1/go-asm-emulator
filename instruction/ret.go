package instruction

import "vinibiavatti1/goasmemulator/machine"

func Ret(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	lineNumber := machine.Stack().Pop()
	machine.IP().Set(lineNumber)
	return nil
}
