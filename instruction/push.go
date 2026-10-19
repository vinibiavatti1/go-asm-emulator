package instruction

import (
	"fmt"

	"github.com/vinibiavatti1/go-asm-emulator/machine"
)

func Push(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	source := instruction.Arg1

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

	err := machine.Stack().Push(value)
	if err != nil {
		return fmt.Errorf("error when pushing to stack: %w", err)
	}

	interpreter.IncInstructionPtr()
	return nil
}
