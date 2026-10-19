package machine

type Register struct {
	value int
}

func NewRegister() *Register {
	return &Register{
		value: 0,
	}
}

func (r *Register) Get() int      { return r.value }
func (r *Register) Set(value int) { r.value = value }
func (r *Register) Inc()          { r.value++ }
func (r *Register) Dec()          { r.value-- }
