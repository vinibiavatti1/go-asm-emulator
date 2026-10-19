package main

import (
	"flag"
	"fmt"
	"os"
	"vinibiavatti1/goasmemulator/instruction"
	"vinibiavatti1/goasmemulator/machine"
)

var src = flag.String("src", "", "ASM source code path")

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

	machine := machine.NewMachine(10, 10)
	interpreter := instruction.NewInterpreter(machine)
	err = interpreter.CompileAndRun(string(code))
	if err != nil {
		panic(err)
	}

	fmt.Printf("ASM (%s) result:\n", path)
	machine.PrintState()
}
