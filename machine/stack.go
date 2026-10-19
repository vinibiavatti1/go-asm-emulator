package machine

import "fmt"

type Stack struct {
	machine *Machine
	data    []int
}

func NewStack(machine *Machine, size int) *Stack {
	if size < 0 {
		size = 0
	}
	return &Stack{
		machine: machine,
		data:    make([]int, size),
	}
}

func (s *Stack) Push(value int) error {
	rsp := s.machine.RSP()
	if rsp.Get() >= len(s.data) {
		return fmt.Errorf("stack overflow. rsp: %d size: %d", rsp.Get(), len(s.data))
	}
	s.data[rsp.Get()] = value
	rsp.Inc()
	return nil
}
func (s *Stack) Pop() int {
	rsp := s.machine.RSP()
	if rsp.Get() > 0 {
		rsp.Dec()
	}
	value := s.data[rsp.Get()]
	s.data[rsp.Get()] = 0
	return value
}
