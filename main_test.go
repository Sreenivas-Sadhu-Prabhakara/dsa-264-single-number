package main

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{-3, 7, -3}, 7},
		{[]int{0, 0, -1}, -1},
	}
	for _, c := range cases {
		if got := SingleNumber(c.in); got != c.want {
			t.Errorf("SingleNumber(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
