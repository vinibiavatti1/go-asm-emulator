package instruction

import (
	"fmt"
	"vinibiavatti1/goasmemulator/machine"
)

func Cmp(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	left := instruction.Arg1
	right := instruction.Arg2

	var y int
	if IsConstant(machine, right) {
		y = MustParseConstant(right)
	} else if IsRegister(machine, right) {
		y = MustGetFromRegister(machine, right)
	} else if IsMemoryAddress(machine, right) {
		y = MustGetFromMemory(machine, right)
	} else {
		return fmt.Errorf("invalid right argument: %s", right)
	}

	var x int
	if IsRegister(machine, left) {
		x = MustGetFromRegister(machine, left)
	} else if IsMemoryAddress(machine, left) {
		x = MustGetFromMemory(machine, left)
	} else {
		return fmt.Errorf("invalid left argument: %s", left)
	}

	flags := ToFlags(x, y)
	machine.EFlags().Set(flags)

	interpreter.IncInstructionPtr()
	return nil
}
