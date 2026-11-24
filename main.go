package main

import "fmt"

// SingleNumber returns the value appearing once when every other value
// appears exactly twice. XOR of everything: pairs annihilate, the loner
// survives.
func SingleNumber(nums []int) int {
	x := 0
	for _, v := range nums {
		x ^= v
	}
	return x
}

func main() {
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2}))
}
