PACKAGE DOCUMENTATION

package pcg
    import "github.com/zeebo/pcg"


TYPES

type T struct {
    // contains filtered or unexported fields
}
    T is a pcg generator.

func New(state uint64) T
    New constructs a pcg with the given state.

func (p *T) Float32() float32
    Float32 returns a float32 uniformly in [0, 1).

func (p *T) Float64() float64
    Float64 returns a float64 uniformly in [0, 1).

func (p *T) Uint32() uint32
    Uint32 returns a random uint32.

func (p *T) Uint32n(n uint32) uint32
    Uint32n returns a uint32 uniformly in [0, n).

func (p *T) Uint64() uint64
    Uint64 returns a random uint64.