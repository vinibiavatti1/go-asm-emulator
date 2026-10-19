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
	fmt.Printf("rax.......: %s (%d)\n", toHexRepr(m.rax.value), m.rax.value)
	fmt.Printf("rbx.......: %s (%d)\n", toHexRepr(m.rbx.value), m.rbx.value)
	fmt.Printf("rcx.......: %s (%d)\n", toHexRepr(m.rcx.value), m.rcx.value)
	fmt.Printf("rdx.......: %s (%d)\n", toHexRepr(m.rdx.value), m.rdx.value)
	fmt.Printf("rsi.......: %s (%d)\n", toHexRepr(m.rsi.value), m.rsi.value)
	fmt.Printf("rdi.......: %s (%d)\n", toHexRepr(m.rdi.value), m.rdi.value)
	fmt.Printf("rsp.......: %s (%d)\n", toHexRepr(m.rsp.value), m.rsp.value)
	fmt.Printf("rbp.......: %s (%d)\n", toHexRepr(m.rbp.value), m.rbp.value)
	fmt.Printf("eflags....: %s (%d)\n", toHexRepr(m.eflags.value), m.eflags.value)

	fmt.Printf("Memory: (%d)\n", len(m.memory.data))
	for address, value := range m.memory.data {
		fmt.Printf("%s: %s (%d)\n", toHexRepr(address), toHexRepr(value), value)
	}

	fmt.Printf("Stack: (%d)\n", len(m.stack.data))
	for address, value := range m.stack.data {
		fmt.Printf("%s: %s (%d)\n", toHexRepr(address), toHexRepr(value), value)
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
