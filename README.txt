PACKAGE DOCUMENTATION

package pcg
    import "github.com/zeebo/pcg"


TYPES

type PT struct {
    // contains filtered or unexported fields
}
    PT is a thread safe pcg generator. The output is non-deterministic, even
    if all of the calls are single threaded. The zero value is valid.

func NewParallel(state uint64) PT
    New constructs a parallel pcg with the given state.

func (p *PT) Float32() float32
    Float32 returns a float32 uniformly in [0, 1). Safe for concurrent
    callers.

func (p *PT) Float64() float64
    Float64 returns a float64 uniformly in [0, 1). Safe for concurrent
    callers.

func (p *PT) Uint32() uint32
    Uint32 returns a random uint32. Safe for concurrent callers.

func (p *PT) Uint32n(n uint32) uint32
    Uint32n returns a uint32 uniformly in [0, n). Safe for concurrent
    callers.

func (p *PT) Uint64() uint64
    Uint64 returns a random uint64. Safe for concurrent callers.

type T struct {
    // contains filtered or unexported fields
}
    T is a pcg generator. The zero value is valid.

func New(state uint64) T
    New constructs a pcg with the given state.

func (p *T) Float32() float32
    Float32 returns a float32 uniformly in [0, 1). Not safe for concurrent
    callers.

func (p *T) Float64() float64
    Float64 returns a float64 uniformly in [0, 1). Not safe for concurrent
    callers.

func (p *T) Uint32() uint32
    Uint32 returns a random uint32. Not safe for concurrent callers.

func (p *T) Uint32n(n uint32) uint32
    Uint32n returns a uint32 uniformly in [0, n). Not safe for concurrent
    callers.

func (p *T) Uint64() uint64
    Uint64 returns a random uint64. Not safe for concurrent callers.
