package pcg

import "testing"

func BenchmarkParPCG(b *testing.B) {
	b.Run("Single", func(b *testing.B) {
		b.Run("Uint32", func(b *testing.B) {
			rng := NewParallel(2345)
			for i := 0; i < b.N; i++ {
				blackholeUint32 += rng.Uint32()
			}
		})

		b.Run("Uint32n", func(b *testing.B) {
			b.Run("Large", func(b *testing.B) {
				rng := NewParallel(2345)
				for i := 0; i < b.N; i++ {
					blackholeUint32 += rng.Uint32n(1<<31 + 1)
				}
			})

			b.Run("Small", func(b *testing.B) {
				rng := NewParallel(2345)
				for i := 0; i < b.N; i++ {
					blackholeUint32 += rng.Uint32n(1000)
				}
			})
		})

		b.Run("Uint64", func(b *testing.B) {
			rng := NewParallel(2345)
			for i := 0; i < b.N; i++ {
				blackholeUint64 += rng.Uint64()
			}
		})

		b.Run("Float64", func(b *testing.B) {
			rng := NewParallel(2345)
			for i := 0; i < b.N; i++ {
				blackholeFloat64 += rng.Float64()
			}
		})

		b.Run("Float32", func(b *testing.B) {
			rng := NewParallel(2345)
			for i := 0; i < b.N; i++ {
				blackholeFloat32 += rng.Float32()
			}
		})
	})

	b.Run("Parallel", func(b *testing.B) {
		b.Run("Uint32", func(b *testing.B) {
			rng := NewParallel(2345)
			b.RunParallel(func(pb *testing.PB) {
				var localUint32 uint32
				for pb.Next() {
					localUint32 += rng.Uint32()
				}
				blackholeUint32 += localUint32
			})
		})

		b.Run("Uint32n", func(b *testing.B) {
			b.Run("Large", func(b *testing.B) {
				rng := NewParallel(2345)
				b.RunParallel(func(pb *testing.PB) {
					var localUint32 uint32
					for pb.Next() {
						localUint32 += rng.Uint32n(1<<31 + 1)
					}
					blackholeUint32 += localUint32
				})
			})

			b.Run("Small", func(b *testing.B) {
				rng := NewParallel(2345)
				b.RunParallel(func(pb *testing.PB) {
					var localUint32 uint32
					for pb.Next() {
						localUint32 += rng.Uint32n(1000)
					}
					blackholeUint32 += localUint32
				})
			})
		})

		b.Run("Uint64", func(b *testing.B) {
			rng := NewParallel(2345)
			b.RunParallel(func(pb *testing.PB) {
				var localUint64 uint64
				for pb.Next() {
					localUint64 += rng.Uint64()
				}
				blackholeUint64 += localUint64
			})
		})

		b.Run("Float64", func(b *testing.B) {
			rng := NewParallel(2345)
			b.RunParallel(func(pb *testing.PB) {
				var localFloat64 float64
				for pb.Next() {
					localFloat64 += rng.Float64()
				}
				blackholeFloat64 += localFloat64
			})
		})

		b.Run("Float32", func(b *testing.B) {
			rng := NewParallel(2345)
			b.RunParallel(func(pb *testing.PB) {
				var localFloat32 float32
				for pb.Next() {
					localFloat32 += rng.Float32()
				}
				blackholeFloat32 += localFloat32
			})
		})
	})
}
