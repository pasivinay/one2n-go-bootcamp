package basics

import (
	"fmt"
	"math"
)

// EvenOdd takes a int and returns whether it's even(true) or odd(false).
func EvenOdd(n int) bool {
	return n%2 == 0
}

// Table takes a int and return a multiplication table for it.
func Table(n int) string {
	table := ""
	for i := 1; i < 11; i++ {
		table = table + fmt.Sprintf("%v x %v = %v\n", n, i, n*i)
		// fmt.Printf("%v x %v = %v",n,i,n*i)
	}

	return table
}

// Sum return sum of n numbers.
func Sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// Swap swaps two numbers from it's variables
func Swap(a, b int) (int, int) {
	c := a
	// swapping numbers
	a = b
	b = c

	return a, b
}

// abs takes two integers and find the number closest to n and divisible by m.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// ClosestNumber takes two intergers n and m, and finds the absolute number closest to n and divisible by m
func ClosestNumber(n, m int) int {
	closest := 0
	min_diff := int(math.MaxInt)

	for i := n - m; i <= n+m; i++ {
		if i%m == 0 {
			diff := abs(n - i)

			if diff < min_diff || (diff == min_diff && abs(i) > abs(closest)) {
				closest = i
				min_diff = diff
			}
		}
	}

	return closest
}

// OppOfDice takes a number on a dice face and return the number on the opposite face.
func OppOfDice(n int) int {
	if n < 1 || n > 7 {
		return -1
	}
	return (7 - n)
}

// NthTermofAP takes two consecutive numbers of an AP and return the nth term of that AP
func NthTermOfAP(a, b, n int) int {
	return (a + (n-1)*(b-a))
}
