package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	got := Iterator("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}

func ExampleIterator() {
    repeat := Iterator("a",4)
    fmt.Println(repeat)
    // Output: aaaa
}