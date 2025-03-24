package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 3)
	want := 5

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
