package instruction

import "github.com/vinibiavatti1/go-asm-emulator/machine"

func Nop(machine *machine.Machine, interpreter *Interpreter, instruction *Instruction) error {
	interpreter.IncInstructionPtr()
	return nil
}
