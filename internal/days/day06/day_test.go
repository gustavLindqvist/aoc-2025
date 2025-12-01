package day06

import "testing"

func TestStar1Example(t *testing.T) {
	result := Star1(Example)
	if result == nil {
		t.Skip("Star1 not implemented")
	}
	t.Logf("Star1 Example: %v", result)
}

func TestStar1(t *testing.T) {
	result := Star1(Input)
	if result == nil {
		t.Skip("Star1 not implemented")
	}
	t.Logf("Star1: %v", result)
}

func TestStar2Example(t *testing.T) {
	result := Star2(Example)
	if result == nil {
		t.Skip("Star2 not implemented")
	}
	t.Logf("Star2 Example: %v", result)
}

func TestStar2(t *testing.T) {
	result := Star2(Input)
	if result == nil {
		t.Skip("Star2 not implemented")
	}
	t.Logf("Star2: %v", result)
}

func BenchmarkStar1(b *testing.B) {
	for b.Loop() {
		Star1(Input)
	}
}

func BenchmarkStar2(b *testing.B) {
	for b.Loop() {
		Star2(Input)
	}
}
