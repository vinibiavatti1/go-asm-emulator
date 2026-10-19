package machine

import "fmt"

type Memory struct {
	data []int
}

func NewMemory(size int) *Memory {
	if size < 0 {
		size = 0
	}
	return &Memory{
		data: make([]int, size),
	}
}

func (m *Memory) Get(address int) (int, error) {
	if address < 0 || address >= len(m.data) {
		return 0, fmt.Errorf("invalid address. address: %d size: %d", address, len(m.data))
	}
	return m.data[address], nil
}
func (m *Memory) Set(address int, value int) error {
	if address < 0 || address >= len(m.data) {
		return fmt.Errorf("invalid address. address: %d size: %d", address, len(m.data))
	}
	m.data[address] = value
	return nil
}
