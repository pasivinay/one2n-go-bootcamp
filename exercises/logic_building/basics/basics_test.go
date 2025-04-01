package basics

import "testing"

func TestEvenOdd(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"Even number", 4, true},
		{"Odd number", 5, false},
		{"Negative even number", -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvenOdd(tt.input)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"Multiplication table of 4", 4, "4 x 1 = 4\n4 x 2 = 8\n4 x 3 = 12\n4 x 4 = 16\n4 x 5 = 20\n4 x 6 = 24\n4 x 7 = 28\n4 x 8 = 32\n4 x 9 = 36\n4 x 10 = 40\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Table(tt.input)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Sum from 1 to 9", 9, 45},
		{"Sum from 1 to 15", 15, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.input)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"Swap 5 and 4", 5, 4, 4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := Swap(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("Expected a: %v, b: %v, got a: %v, b: %v", tt.wantA, tt.wantB, gotA, gotB)
			}
		})
	}
}

func TestClosestNumber(t *testing.T) {
	tests := []struct {
		name     string
		input, d int
		want     int
	}{
		{"Closest to 13 divisible by 4", 13, 4, 12},
		{"Closest to -15 divisible by 6", -15, 6, -18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClosestNumber(tt.input, tt.d)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestOppOfDice(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Opposite of dice face 5", 5, 2},
		{"Opposite of dice face 3", 3, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OppOfDice(tt.input)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNthTermOfAP(t *testing.T) {
	tests := []struct {
		name           string
		first, diff, n int
		want           int
	}{
		{"Nth term of AP (1,2,6)", 1, 2, 6, 6},
		{"Nth term of AP (2,5,6)", 2, 5, 6, 17},
		{"Nth term of AP (-2,-4,8)", -2, -4, 8, -16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NthTermOfAP(tt.first, tt.diff, tt.n)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
