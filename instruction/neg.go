package instruction

import (
	"fmt"
	"vinibiavatti1/goasmemulator/machine"
)

func Neg(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	source := instruction.Arg1

	if IsRegister(machine, source) {
		x := MustGetFromRegister(machine, source)
		MustSetInRegister(machine, source, -x)
	} else if IsMemoryAddress(machine, source) {
		x := MustGetFromMemory(machine, source)
		MustSetInMemory(machine, source, -x)
	} else {
		return fmt.Errorf("invalid source: %s", source)
	}

	interpreter.IncInstructionPtr()
	return nil
}
