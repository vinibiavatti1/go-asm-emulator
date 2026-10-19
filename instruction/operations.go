package instruction

import (
	"strconv"
	"strings"
	"vinibiavatti1/goasmemulator/machine"
)

func IsRegister(machine *machine.Machine, registerName string) (ok bool) {
	_, ok = machine.Register(registerName)
	return
}

func IsMemoryAddress(machine *machine.Machine, registerName string) bool {
	if !strings.HasPrefix(registerName, "[") || !strings.HasSuffix(registerName, "]") {
		return false
	}
	registerName = strings.TrimLeft(registerName, "[")
	registerName = strings.TrimRight(registerName, "]")
	return IsRegister(machine, registerName)
}

func IsConstant(machine *machine.Machine, value string) bool {
	if strings.HasPrefix(value, "0x") {
		value = strings.TrimPrefix(value, "0x")
		_, err := strconv.ParseInt(value, 16, 64)
		return err == nil
	} else if strings.HasPrefix(value, "0b") {
		value = strings.TrimPrefix(value, "0b")
		_, err := strconv.ParseInt(value, 2, 64)
		return err == nil
	} else {
		_, err := strconv.ParseInt(value, 10, 64)
		return err == nil
	}
}

func MustParseConstant(value string) int {
	var result int64
	var err error
	if strings.HasPrefix(value, "0x") {
		value = strings.TrimPrefix(value, "0x")
		result, err = strconv.ParseInt(value, 16, 64)
	} else if strings.HasPrefix(value, "0b") {
		value = strings.TrimPrefix(value, "0b")
		result, err = strconv.ParseInt(value, 2, 64)
	} else {
		result, err = strconv.ParseInt(value, 10, 64)
	}
	if err != nil {
		panic("constant could not be parsed")
	}
	return int(result)
}

func MustGetFromRegister(machine *machine.Machine, registerName string) int {
	register, ok := machine.Register(registerName)
	if !ok {
		panic("register not found: " + registerName)
	}
	return register.Get()
}

func MustSetInRegister(machine *machine.Machine, registerName string, value int) {
	register, ok := machine.Register(registerName)
	if !ok {
		panic("register not found: " + registerName)
	}
	register.Set(value)
}

func MustGetFromMemory(machine *machine.Machine, registerName string) int {
	registerName = strings.TrimLeft(registerName, "[")
	registerName = strings.TrimRight(registerName, "]")
	register, ok := machine.Register(registerName)
	if !ok {
		panic("register not found: " + registerName)
	}
	address := register.Get()
	value, err := machine.Memory().Get(address)
	if err != nil {
		panic("could not get from memory")
	}
	return value
}

func MustSetInMemory(machine *machine.Machine, registerName string, value int) {
	registerName = strings.TrimLeft(registerName, "[")
	registerName = strings.TrimRight(registerName, "]")
	register, ok := machine.Register(registerName)
	if !ok {
		panic("register not found: " + registerName)
	}
	address := register.Get()
	err := machine.Memory().Set(address, value)
	if err != nil {
		panic("could not set in memory")
	}
}

func ToFlags(x int, y int) int {
	result := x - y
	zf := result == 0
	sf := result < 0
	flags := 0
	if zf {
		flags = flags | 1
	}
	if sf {
		flags = flags | 2
	}
	return flags
}

func FromFlags(flags int) (zf bool, sf bool) {
	zf = flags&1 == 1
	sf = flags&2 == 2
	return
}
