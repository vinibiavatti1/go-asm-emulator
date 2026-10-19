package instruction

import (
	"fmt"

	"github.com/vinibiavatti1/goasmemulator/machine"
)

func Mov(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	destination := instruction.Arg1
	source := instruction.Arg2

	var value int
	if IsConstant(machine, source) {
		value = MustParseConstant(source)
	} else if IsRegister(machine, source) {
		value = MustGetFromRegister(machine, source)
	} else if IsMemoryAddress(machine, source) {
		value = MustGetFromMemory(machine, source)
	} else {
		return fmt.Errorf("invalid source: %s", source)
	}

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
