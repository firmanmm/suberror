package suberror

import "testing"

func BenchmarkDeriveErrorDepth10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parent := RuntimeError.Derive()
		for j := 0; j < 10; j++ {
			parent = parent.Derive()
		}
	}
}

func BenchmarkDeriveErrorDepth25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parent := RuntimeError.Derive()
		for j := 0; j < 25; j++ {
			parent = parent.Derive()
		}
	}
}

func BenchmarkDeriveErrorDepth50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parent := RuntimeError.Derive()
		for j := 0; j < 50; j++ {
			parent = parent.Derive()
		}
	}
}

func BenchmarkDeriveErrorDepth100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parent := RuntimeError.Derive()
		for j := 0; j < 100; j++ {
			parent = parent.Derive()
		}
	}
}

func BenchmarkCheckTypeErrorDepth100TargetRoot(b *testing.B) {
	target := RuntimeError.Derive()
	child := target
	for j := 0; j < 100; j++ {
		child = child.Derive()
	}
	for i := 0; i < b.N; i++ {
		if !child.TypeOf(target) {
			b.Errorf("Current type is not parent")
		}
	}
}

func BenchmarkCheckTypeErrorDepth100Target50(b *testing.B) {
	target := RuntimeError.Derive()
	child := target
	for j := 0; j < 100; j++ {
		child = child.Derive()
		if j == 50 {
			target = child
		}
	}
	for i := 0; i < b.N; i++ {
		if !child.TypeOf(target) {
			b.Errorf("Current type is not parent")
		}
	}
}

func BenchmarkCheckTypeErrorDepth100TargetLeaf(b *testing.B) {
	target := RuntimeError.Derive()
	for j := 0; j < 100; j++ {
		target = target.Derive()
	}
	for i := 0; i < b.N; i++ {
		if !target.TypeOf(target) {
			b.Errorf("Current type is not parent")
		}
	}
}
