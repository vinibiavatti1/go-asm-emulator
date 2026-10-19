package instruction

import (
	"fmt"
	"slices"
	"strings"

	"github.com/vinibiavatti1/go-asm-emulator/machine"
)

type Interpreter struct {
	machine      *machine.Machine
	instructions []*Instruction
	labels       map[string]int
}

func NewInterpreter(machine *machine.Machine) *Interpreter {
	return &Interpreter{
		machine:      machine,
		instructions: []*Instruction{},
		labels:       map[string]int{},
	}
}

func (i *Interpreter) Instructions() []*Instruction { return i.instructions }
func (i *Interpreter) Labels() map[string]int       { return i.labels }
func (i *Interpreter) IncInstructionPtr()           { i.machine.IncInstructionPtr() }
func (i *Interpreter) CompileAndRun(code string) error {
	i.labels = map[string]int{}
	i.instructions = []*Instruction{}
	i.machine.IP().Set(0)
	err := i.compile(code)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}
	err = i.run()
	if err != nil {
		return fmt.Errorf("runtime error: %w", err)
	}
	return nil
}
func (i *Interpreter) run() error {
	for {
		ip := i.machine.IP().Get()
		if ip < 0 || ip >= len(i.instructions) {
			return nil
		}
		instruction := i.instructions[ip]
		err := EvaluateInstruction(i.machine, i, instruction)
		if err != nil {
			return fmt.Errorf("Instruction %s on line %d raised an error: %w", instruction.Name, instruction.LineNumber, err)
		}
	}
}
func (i *Interpreter) compile(code string) error {
	for lineNumber, line := range strings.Split(code, "\n") {
		err := i.compileLine(lineNumber, line)
		if err != nil {
			return fmt.Errorf("failed to compile code: %w", err)
		}
	}
	return nil
}
func (i *Interpreter) compileLine(lineNumber int, line string) error {
	// Remove comments
	if strings.Contains(line, ";") {
		line = strings.Split(line, ";")[0]
	}

	// Remove commas
	if strings.Contains(line, ",") {
		line = strings.ReplaceAll(line, ",", "")
	}

	// Normalize
	line = strings.Trim(line, " ")
	line = strings.ReplaceAll(line, "\t", " ")
	line = strings.ToLower(line)

	// Get entities
	var label string
	var name string
	var arg1 string
	var arg2 string
	fields := strings.Fields(line)
	if len(fields) == 0 {
		instruction := NewNopInstruction(lineNumber)
		i.instructions = append(i.instructions, instruction)
		return nil
	}
	if !strings.HasSuffix(fields[0], ":") {
		fields = slices.Insert(fields, 0, "")
	}
	for part := len(fields); part < 4; part++ {
		fields = append(fields, "")
	}
	label = fields[0]
	name = fields[1]
	arg1 = fields[2]
	arg2 = fields[3]

	// Set empty instruction to nop
	if name == "" {
		name = "nop"
	}

	// Map label
	label = strings.Replace(label, ":", "", 1)
	if label != "" {
		i.labels[label] = lineNumber
	}

	// Build instruction
	instruction := NewInstruction(lineNumber, label, name, arg1, arg2)
	i.instructions = append(i.instructions, instruction)
	return nil
}
func (i *Interpreter) JumpToLabel(label string) error {
	lineNumber, ok := i.labels[label]
	if !ok {
		return fmt.Errorf("label %s not found", label)
	}
	i.machine.IP().Set(lineNumber)
	return nil
}
