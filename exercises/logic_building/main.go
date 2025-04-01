package main

import (
	"fmt"

	"one2n.io/golang-logic-building/basics"
)

func main() {
	fmt.Print(basics.EvenOdd(5), "\n")
	fmt.Print(basics.Table(4), "\n")
	fmt.Print(basics.Sum(15), "\n")
	fmt.Print(basics.Swap(5, 8))
	fmt.Print(basics.ClosestNumber(-15, 6), "\n")
	fmt.Print(basics.OppOfDice(1), "\n")
}
