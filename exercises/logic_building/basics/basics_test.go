package basics

import (
	"testing"
)

func TestEvenOdd(t *testing.T) {
	t.Run("Check whether the 4 is even or odd", func(t *testing.T) {
		got := EvenOdd(4)
		want := true

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
	t.Run("Check whether the 5 is even or odd", func(t *testing.T) {
		got := EvenOdd(5)
		want := false

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
	t.Run("Check whether the -10 is even or odd", func(t *testing.T) {
		got := EvenOdd(-10)
		want := true

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
}

func TestTable(t *testing.T) {
	t.Run("Get the table of 4", func(t *testing.T) {
		got := Table(4)
		want := "4 x 1 = 4\n4 x 2 = 8\n4 x 3 = 12\n4 x 4 = 16\n4 x 5 = 20\n4 x 6 = 24\n4 x 7 = 28\n4 x 8 = 32\n4 x 9 = 36\n4 x 10 = 40\n"

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("Get the sum of numbers 1 to 9", func(t *testing.T) {
		got := Sum(9)
		want := 45

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
	t.Run("Get the sum of numbers 1 to 15", func(t *testing.T) {
		got := Sum(15)
		want := 120

		if got != want {
			t.Errorf("The answer is %v but we got %v", want, got)
		}
	})
}

func TestSwap(t *testing.T) {
	t.Run("Swap the numbers", func(t *testing.T) {
		a, b := Swap(5, 4)
		c, d := 4, 5

		if b != d && a != c {
			t.Errorf("Expected a: %v b: %v, got a: %v b: %v", c, d, a, b)
		}
	})
}

func TestClosestNumber(t *testing.T) {
	t.Run("Get the number closest to 13 and divisible by 4", func(t *testing.T) {
		got := ClosestNumber(13, 4)
		want := 12

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
	t.Run("Get the number closest to -15 and divisible by 6", func(t *testing.T) {
		got := ClosestNumber(-15, 6)
		want := -18

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
}

func TestOppOfDice(t *testing.T) {
	t.Run("Get the number on the opposite side of the dice with face number 5", func(t *testing.T) {
		got := OppOfDice(5)
		want := 2

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
	t.Run("Get the number on the opposite side of the dice with face number 3", func(t *testing.T) {
		got := OppOfDice(3)
		want := 4

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
}

func TestNthTermOfAP(t *testing.T) {
    t.Run("Get the number on the opposite side of the dice with face number 3", func(t *testing.T) {
		got := NthTermOfAP(1,2,6) // 1, 2, 3, 4, 5 ,6
		want := 6

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
    t.Run("Get the number on the opposite side of the dice with face number 3", func(t *testing.T) {
		got := NthTermOfAP(2,5,6) //2, 5, 8, 11, 14, 17
		want := 17

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
    t.Run("Get the number on the opposite side of the dice with face number 3", func(t *testing.T) {
		got := NthTermOfAP(-2,-4,8) //-2, -4, -6, -8, -10, -12, -14, -16
		want := -16

		if got != want {
			t.Errorf("Expected : %v , got : %v", want, got)
		}
	})
}
