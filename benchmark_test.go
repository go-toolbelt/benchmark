package benchmark_test

import (
	"reflect"
	"testing"

	"github.com/go-toolbelt/benchmark"
)

func TestZeroAllocsPerOp(t *testing.T) {
	benchmark.Test(
		t,
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
			}
		},
		benchmark.ZeroAllocsPerOp(),
	)
}

func TestOneAllocPerOp(t *testing.T) {
	benchmark.Test(
		t,
		func(b *testing.B) {
			alloced := make([]*int64, 0, b.N)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				i64 := int64(i)
				alloced = append(alloced, &i64)
			}
			if len(alloced) != b.N {
				t.Error("Unexpected length to slice alloced")
			}
		},
		benchmark.AllocsPerOp(1),
		benchmark.AllocedBytesPerOp(int64(reflect.TypeOf(int64(0)).Size())),
	)
}
