package main

import (
	"fmt"

	"one2n.io/golang-logic-building/basics"
)

func main() {
	fmt.Print(basics.EvenOdd(),"\n")
	fmt.Print(basics.GetTable(),"\n")
	fmt.Print(basics.Sum(),"\n")
	basics.Swap()
}
