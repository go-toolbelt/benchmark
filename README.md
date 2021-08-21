# benchmark
Simple Go utility for writing benchmark tests

# Example Usage

```go
func TestBenchmarkExample(t *testing.T) {
	benchmark.Test(
		t,
		func(b *testing.B) {
			// Write code to benchmark here.
			...
		},
		benchmark.ZeroAllocsPerOp(),
	)
}
```
