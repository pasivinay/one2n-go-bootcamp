package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
    got := Sum([]int{3,2,4})
    want := 9

    if got != want {
        t.Errorf("Expected %v but got %v", want, got)
    }
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 3, 4})
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}