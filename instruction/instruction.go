package instruction

type Instruction struct {
	LineNumber int
	Label      string
	Name       string
	Arg1       string
	Arg2       string
}

func NewInstruction(lineNumber int, label string, name string, arg1 string, arg2 string) *Instruction {
	return &Instruction{
		LineNumber: lineNumber,
		Label:      label,
		Name:       name,
		Arg1:       arg1,
		Arg2:       arg2,
	}
}

func NewNopInstruction(lineNumber int) *Instruction {
	return NewInstruction(lineNumber, "", "nop", "", "")
}
