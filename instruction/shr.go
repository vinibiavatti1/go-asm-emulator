package instruction

import (
	"fmt"

	"github.com/vinibiavatti1/go-asm-emulator/machine"
)

func Shr(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	destination := instruction.Arg1
	source := instruction.Arg2

	var y int
	if IsConstant(machine, source) {
		y = MustParseConstant(source)
	} else if IsRegister(machine, source) {
		y = MustGetFromRegister(machine, source)
	} else if IsMemoryAddress(machine, source) {
		y = MustGetFromMemory(machine, source)
	} else {
		return fmt.Errorf("invalid source: %s", source)
	}

	var x int
	if IsRegister(machine, destination) {
		x = MustGetFromRegister(machine, destination)
		MustSetInRegister(machine, destination, x>>y)
	} else if IsMemoryAddress(machine, destination) {
		x = MustGetFromMemory(machine, destination)
		MustSetInMemory(machine, destination, x>>y)
	} else {
		return fmt.Errorf("invalid destination: %s", source)
	}

	interpreter.IncInstructionPtr()
	return nil
}
