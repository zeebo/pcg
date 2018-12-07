package pcg

// T is a pcg generator.
type T struct{ state uint64 }

// mul is the multiplier for the LCG step.
const (
	mul = 6364136223846793005
	inc = 11981177638785157926
)

// New constructs a pcg with the given state.
func New(state uint64) T { return T{state} }

// next advances and returns the state.
func (p *T) next() uint64 {
	p.state = p.state*mul + inc
	return p.state
}

// Uint32 returns a random uint32.
func (p *T) Uint32() uint32 {
	state := p.next()

	xor := uint32(((state >> 18) ^ state) >> 27)
	shift := uint(state>>59) & 31

	return xor>>shift | xor<<(32-shift)
}

// Uint32n returns a uint32 uniformly in [0, n).
func (p *T) Uint32n(n uint32) uint32 {
	if n == 0 {
		return 0
	}

	x := p.Uint32()
	m := uint64(x) * uint64(n)
	l := uint32(m)

	if l < n {
		t := -n
		if t >= n {
			t -= n
			if t >= n {
				t = t % n
			}
		}

	again:
		if l < t {
			x = p.Uint32()
			m = uint64(x) * uint64(n)
			l = uint32(m)
			goto again
		}
	}

	return uint32(m >> 32)
}

// Uint64 returns a random uint64.
func (p *T) Uint64() uint64 {
	state1 := p.next()
	state2 := p.next()

	xor1 := uint32(((state1 >> 18) ^ state1) >> 27)
	shift1 := uint(state1>>59) & 31

	xor2 := uint32(((state2 >> 18) ^ state2) >> 27)
	shift2 := uint(state2>>59) & 31

	return uint64(xor1>>shift1|xor1<<(32-shift1))<<32 |
		uint64(xor2>>shift2|xor2<<(32-shift2))
}

// Float64 returns a float64 uniformly in [0, 1).
func (p *T) Float64() float64 {
again:
	state1 := p.next()
	state2 := p.next()

	xor1 := uint32(((state1 >> 18) ^ state1) >> 27)
	shift1 := uint(state1>>59) & 31

	xor2 := uint32(((state2 >> 18) ^ state2) >> 27)
	shift2 := uint(state2>>59) & 31

	v := uint64(xor1>>shift1|xor1<<(32-shift1)) |
		uint64(xor2>>shift2|xor2<<(32-shift2))

	out := float64(v>>(64-53)) / (1 << 53)
	if out == 1 {
		goto again
	}

	return out
}

// Float32 returns a float32 uniformly in [0, 1).
func (p *T) Float32() float32 {
again:
	out := float32(p.Uint32()>>(32-24)) / (1 << 24)
	if out == 1 {
		goto again
	}

	return out
}
