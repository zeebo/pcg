package pcg

// T is a pcg generator.
type T struct {
	State uint64
	Inc   uint64
}

// mul is the multiplier for the LCG step.
const mul = 6364136223846793005

// New constructs a pcg with the given state and inc.
func New(state, inc uint64) T {
	inc = inc<<1 | 1
	return T{
		State: (inc+state)*mul + inc,
		Inc:   inc,
	}
}

// Uint32 returns a random uint32.
func (p *T) Uint32() uint32 {
	if p.Inc == 0 {
		*p = T{mul + 1, 1}
	}

	state := p.State
	p.State = state*mul + p.Inc

	xor := uint32(((state >> 18) ^ state) >> 27)
	shift := uint(state>>59) & 31

	return xor>>shift | xor<<(32-shift)
}

// Uint64 advances twice and returns a random uint64.
func (p *T) Uint64() uint64 {
	state1 := p.State
	state2 := state1*mul + p.Inc
	p.State = state2*mul + p.Inc

	xor1 := uint32(((state1 >> 18) ^ state1) >> 27)
	shift1 := uint(state1>>59) & 31

	xor2 := uint32(((state2 >> 18) ^ state2) >> 27)
	shift2 := uint(state2>>59) & 31

	return uint64(xor1>>shift1|xor1<<(32-shift1))<<32 |
		uint64(xor2>>shift2|xor2<<(32-shift2))
}

// Intn returns an int uniformly in [0, n)
func (p *T) Intn(n int) int {
	return int((uint64(p.Uint32()) * uint64(n)) >> 32)
}

// Float64 returns a float uniformly in [0, 1)
func (p *T) Float64() float64 {
	state1 := p.State
	state2 := state1*mul + p.Inc
	p.State = state2*mul + p.Inc

	xor1 := uint32(((state1 >> 18) ^ state1) >> 27)
	shift1 := uint(state1>>59) & 31

	xor2 := uint32(((state2 >> 18) ^ state2) >> 27)
	shift2 := uint(state2>>59) & 31

	v := uint64(xor1>>shift1|xor1<<(32-shift1)) |
		uint64(xor2>>shift2|xor2<<(32-shift2))

	return float64(v>>(64-53)) / (1 << 53)
}
