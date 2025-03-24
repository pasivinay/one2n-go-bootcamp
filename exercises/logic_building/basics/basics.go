package basics

import (
	"fmt"
	_ "math"
)

// This function takes a int and returns whether it's even(true) or odd(false).

func EvenOdd() bool {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)
	// fmt.Printf("%v",n%2==0)
	return n%2 == 0
}

// This function takes a int and return a multiplication table for it.
func GetTable() string {
	var n int
	fmt.Print("Enter a number for a table :")
	fmt.Scanf("%v", &n)
	table := ""
	for i := 1; i < 11; i++ {
		table = table + fmt.Sprintf("%v x %v = %v\n", n, i, n*i)
		// fmt.Printf("%v x %v = %v",n,i,n*i)
	}

	return table
}

// This function return sum of n numbers.

func Sum() int {
	var n int
	fmt.Print("Enter a number for sum of n numbers :")
	fmt.Scanf("%v", &n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

// This function swaps two numbers from it's variables

func Swap() {
	var a, b, c int

	fmt.Println("Input:")
	fmt.Scanf("%v %v", &a, &b)

	fmt.Println("Numbers before swapping: \n a =", a, "\n b =", b)

	c = a
	a = b
	b = c

	fmt.Println("Numbers after swapping:\n a =", a, "\n b =", b)
}

// This function takes 

// func ClosestNumber() {
// 	closest := math.MaxInt
// 	var n, m int
// 	fmt.Scanf("%v %v", &n, &m)
// 	for i := n - m; i <= n+m; i++ {
// 		fmt.Println(closest, i, i%m)
// 		if i%m == 0 && (n-i) >= (n-closest) && math.Abs(float64(closest)) >= math.Abs(float64(closest)) {
// 			closest = i
// 			fmt.Println(closest, i)
// 		}
// 	}
// }
