/**
Ferrun 2018
You are given an array A consisting of the integers −1, 0 and 1. A slice of that array is any pair of integers (P, Q) such that 0 ≤ P ≤ Q < N. Your task is to find the longest slice of A whose elements yield a non-negative sum.

Write a function:

func Solution(A []int) int

that, given an array A of length N, consisting only of the values −1, 0, 1, returns the length of the longest slice of A that yields a non-negative sum. If there's no such slice, your function should return 0.

For example, given A = [−1, −1, 1, −1, 1, 0, 1, −1, −1], your function should return 7, as the slice starting at the second position and ending at the eighth is the longest slice with a non-negative sum.

For another example, given A = [1, 1, −1, −1, −1, −1, −1, 1, 1] your function should return 4: both the first four elements and the last four elements of array A are longest valid slices.

Assume that:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1..1].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
**/

/*
Result: 90%! One test case is mysteriously failing
https://app.codility.com/demo/results/training9Y33RA-GD3/
*/

package main

// package solution

// you can also use imports, for example:
import "fmt"

// import "os"

import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {

	N := len(A)

	sum := make([]int, N)
	sum[0] = A[0]
	for i := 1; i < N; i++ {
		sum[i] = sum[i-1] + A[i]
	}

	if sum[N-1] >= 0 {
		return N
	}

	// Now lets look for repitions in sum, basically the next position where a sum repeats
	// is indicative of non zero (especially for negative values)
	lower := make(map[int]int)
	upper := make(map[int]int)
	for i := 0; i < N; i++ {
		s := sum[i]
		_, ok := lower[s]
		if !ok {
			lower[s] = i
			upper[s] = -1 // dummy
		} else {
			upper[s] = i
		}
	}

	// fmt.Printf("lower: %v\n", lower)
	// fmt.Printf("upper: %v\n", upper)

	var sum_arr []int
	for s, _ := range lower {
		sum_arr = append(sum_arr, s)
	}

	sort.Ints(sum_arr)

	max_gap := 0
	first_lower := N
	for i := 0; i < len(sum_arr); i++ {
		s := sum_arr[i]
		l := lower[s]
		if l < first_lower {
			first_lower = l
			//fmt.Printf("first_lower: %v\n", first_lower)
		}

		u := upper[s]
		// fmt.Printf("u: %v\n", u)
		gap := u - first_lower
		if s >= 0 {
			// or 0 or more the values are inclusive by 1
			gap++
		}
		if gap > max_gap {
			max_gap = gap
		}
	}

	return max_gap
}

func main() {
	var a1 = []int{0, -1, 0, 0, 1, 0, -1, -1}
	ret := Solution(a1)
	fmt.Printf("ret: %v\n", ret)
}
