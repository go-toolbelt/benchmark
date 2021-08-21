package benchmark

import "testing"

type Assertion func(t *testing.T, result testing.BenchmarkResult)

func Test(
	t *testing.T,
	f func(b *testing.B),
	assertions ...Assertion,
) {
	result := testing.Benchmark(f)
	for _, assertion := range assertions {
		assertion(t, result)
	}
}

func ZeroAllocsPerOp() Assertion {
	return func(t *testing.T, result testing.BenchmarkResult) {
		AllocsPerOp(0)(t, result)
		AllocedBytesPerOp(0)(t, result)
	}
}

func AllocsPerOp(expected int64) Assertion {
	return func(t *testing.T, result testing.BenchmarkResult) {
		actual := result.AllocsPerOp()
		if expected != actual {
			t.Errorf("unexpected allocs per op: expected %d got %d", expected, actual)
		}
	}
}

func AllocedBytesPerOp(expected int64) Assertion {
	return func(t *testing.T, result testing.BenchmarkResult) {
		actual := result.AllocedBytesPerOp()
		if expected != actual {
			t.Errorf("unexpected alloced bytes per op: expected %d got %d", expected, actual)
		}
	}
}
