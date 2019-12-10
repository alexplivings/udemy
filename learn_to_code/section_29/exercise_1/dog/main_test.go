package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(10)

	if x != 70 {
		t.Errorf("Expected 70, got %v", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(10)

	if x != 70 {
		t.Errorf("Expected 70, got %v", x)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
