PACKAGE DOCUMENTATION

package pcg
    import "."


TYPES

type T struct {
    State uint64
    Inc   uint64
}
    T is a pcg generator. The zero value is invalid.

func New(state, inc uint64) T
    New constructs a pcg with the given state and inc.

func (p *T) Float64() float64
    Float64 returns a float uniformly in [0, 1)

func (p *T) Intn(n int) int
    Intn returns an int uniformly in [0, n)

func (p *T) Uint32() uint32
    Uint32 returns a random uint32.


