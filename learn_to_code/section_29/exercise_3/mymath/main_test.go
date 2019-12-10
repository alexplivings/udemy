package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 4},
		[]int{4, 6, 8, 100},
		[]int{80, 90, 95, 100, 110},
	}

	answers := []float64{2.5, 7, 95}

	for i, v := range tests {
		x := CenteredAvg(v)

		if answers[i] != x {
			t.Errorf("Expected %v, got %v", answers[i], x)
		}
	}
}

func ExampleCenteredAvg() {
	test := []int{1, 2, 3, 4}
	fmt.Println(CenteredAvg(test))
	// Output:
	// 2.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	test := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		CenteredAvg(test)
	}
}
