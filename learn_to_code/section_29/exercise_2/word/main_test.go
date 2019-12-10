package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	testString := "Hello this is a test"
	wc := Count(testString)

	if wc != 5 {
		t.Errorf("Expected 5, got %v", wc)
	}
}

func TestUseCount(t *testing.T) {
	testString := "Another test test test string"
	wc := UseCount(testString)

	if wc["test"] != 3 {
		t.Errorf("Expected 3, got %v", wc["test"])
	}
}

func ExampleCount() {
	testString := "Big test string thing"
	fmt.Println(Count(testString))
	// Output:
	// 4
}

func ExampleUseCount() {
	testString := "Big test test thing"
	testOutput := UseCount(testString)
	fmt.Println(testOutput["test"])
	// Output:
	// 2
}

func BenchmarkCount(b *testing.B) {
	testString := "Hello this is a test"
	for i := 0; i < b.N; i++ {
		Count(testString)
	}
}

func BenchmarkUseCount(b *testing.B) {
	testString := "Hello this is a test"
	for i := 0; i < b.N; i++ {
		UseCount(testString)
	}
}
