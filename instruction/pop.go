package instruction

import (
	"fmt"

	"github.com/vinibiavatti1/goasmemulator/machine"
)

func Pop(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	destination := instruction.Arg1

	value := machine.Stack().Pop()
	if IsRegister(machine, destination) {
		MustSetInRegister(machine, destination, value)
	} else if IsMemoryAddress(machine, destination) {
		MustSetInMemory(machine, destination, value)
	} else {
		return fmt.Errorf("invalid destination: %s", destination)
	}

	interpreter.IncInstructionPtr()
	return nil
}
