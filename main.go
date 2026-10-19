package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vinibiavatti1/go-asm-emulator/instruction"
	"github.com/vinibiavatti1/go-asm-emulator/machine"
)

var src = flag.String("src", "", "ASM source code path")
var memSize = flag.Int("memory", 8, "Memory size")
var stackSize = flag.Int("stack", 8, "Stack size")

func init() {
	flag.Parse()
}

func main() {
	path := *src
	if path == "" {
		panic("the --src flag must be specified")
	}

	code, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("error to read %s: %w", path, err)
		panic(err)
	}

	machine := machine.NewMachine(*memSize, *stackSize)
	interpreter := instruction.NewInterpreter(machine)
	err = interpreter.CompileAndRun(string(code))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result (%s):\n", path)
	machine.PrintState()
}
