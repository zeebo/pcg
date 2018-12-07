package pcg

import "testing"

func BenchmarkTID(b *testing.B) {
	b.Run("Single", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			tid()
		}
	})

	b.Run("Parallel", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				tid()
			}
		})
	})

}
