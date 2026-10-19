package machine

import (
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	rax    *Register
	rbx    *Register
	rcx    *Register
	rdx    *Register
	rsi    *Register
	rdi    *Register
	rsp    *Register
	rbp    *Register
	eflags *Register
	ip     *Register
	memory *Memory
	stack  *Stack
}

func NewMachine(memorySize int, stackSize int) *Machine {
	machine := &Machine{
		rax:    NewRegister(),
		rbx:    NewRegister(),
		rcx:    NewRegister(),
		rdx:    NewRegister(),
		rsi:    NewRegister(),
		rdi:    NewRegister(),
		rsp:    NewRegister(),
		rbp:    NewRegister(),
		eflags: NewRegister(),
		ip:     NewRegister(),
		memory: NewMemory(memorySize),
		stack:  nil,
	}
	machine.stack = NewStack(machine, stackSize)
	return machine
}

func (m *Machine) RAX() *Register     { return m.rax }
func (m *Machine) RBX() *Register     { return m.rbx }
func (m *Machine) RCX() *Register     { return m.rcx }
func (m *Machine) RDX() *Register     { return m.rdx }
func (m *Machine) RSI() *Register     { return m.rsi }
func (m *Machine) RDI() *Register     { return m.rdi }
func (m *Machine) RSP() *Register     { return m.rsp }
func (m *Machine) RBP() *Register     { return m.rbp }
func (m *Machine) EFlags() *Register  { return m.eflags }
func (m *Machine) IP() *Register      { return m.ip }
func (m *Machine) IncInstructionPtr() { m.IP().Inc() }
func (m *Machine) Memory() *Memory    { return m.memory }
func (m *Machine) Stack() *Stack      { return m.stack }
func (m *Machine) PrintState() {
	fmt.Println("Registers:")
	fmt.Printf("rax: %s\n", toHexRepr(m.rax.value))
	fmt.Printf("rbx: %s\n", toHexRepr(m.rbx.value))
	fmt.Printf("rcx: %s\n", toHexRepr(m.rcx.value))
	fmt.Printf("rdx: %s\n", toHexRepr(m.rdx.value))
	fmt.Printf("rsi: %s\n", toHexRepr(m.rsi.value))
	fmt.Printf("rdi: %s\n", toHexRepr(m.rdi.value))
	fmt.Printf("rsp: %s\n", toHexRepr(m.rsp.value))
	fmt.Printf("rbp: %s\n", toHexRepr(m.rbp.value))
	fmt.Printf("eflags: %s\n", toHexRepr(m.eflags.value))

	fmt.Println("Memory:")
	for address, value := range m.memory.data {
		fmt.Printf("%s: %s\n", toHexRepr(address), toHexRepr(value))
	}

	fmt.Println("Stack:")
	for address, value := range m.stack.data {
		fmt.Printf("%s: %s\n", toHexRepr(address), toHexRepr(value))
	}
}
func (m *Machine) Register(name string) (*Register, bool) {
	switch name {
	case "rax":
		return m.rax, true
	case "rbx":
		return m.rbx, true
	case "rcx":
		return m.rcx, true
	case "rdx":
		return m.rdx, true
	case "rsi":
		return m.rsi, true
	case "rdi":
		return m.rdi, true
	case "rsp":
		return m.rsp, true
	case "rbp":
		return m.rbp, true
	default:
		return nil, false
	}
}

func toHexRepr(value int) (result string) {
	result = strconv.FormatInt(int64(value), 16)
	result = strings.ToUpper(result)
	for i := len(result); i < 8; i++ {
		result = "0" + result
	}
	result = "0x" + result
	return
}
