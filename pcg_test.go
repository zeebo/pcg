package pcg

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestPCG(t *testing.T) {
	t.Run("Matches", func(t *testing.T) {
		rng := New(2345, 2378)
		out := make([]uint32, 10)
		for i := range out {
			out[i] = rng.Uint32()
		}

		assert.DeepEqual(t, out, []uint32{
			0xccca066b,
			0x40cee775,
			0x0df46902,
			0x981fbe29,
			0xfc8bfb85,
			0xcfd9eef2,
			0xa046c325,
			0x31abe14c,
			0xe29defb4,
			0x160568cc,
		})
	})

	t.Run("Zero", func(t *testing.T) {
		var rng1 T
		var rng2 = New(0, 0)

		for i := 0; i < 10; i++ {
			assert.Equal(t, rng1.Uint32(), rng2.Uint32())
		}
	})
}

var (
	blackholeUint32  uint32
	blackholeUint64  uint64
	blackholeFloat64 float64
)

func BenchmarkPCG(b *testing.B) {
	b.Run("Uint32", func(b *testing.B) {
		rng := New(2345, 2378)
		for i := 0; i < b.N; i++ {
			blackholeUint32 += rng.Uint32()
		}
	})

	b.Run("Uint64", func(b *testing.B) {
		rng := New(2345, 2378)
		for i := 0; i < b.N; i++ {
			blackholeUint64 += rng.Uint64()
		}
	})

	b.Run("Float64", func(b *testing.B) {
		rng := New(2345, 2378)
		for i := 0; i < b.N; i++ {
			blackholeFloat64 += rng.Float64()
		}
	})
}
