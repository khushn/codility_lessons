/**
You are playing a game with N cards. On both sides of each card there is a positive integer. The cards are laid on the table. The score of the game is the smallest positive integer that does not occur on the face-up cards. You may flip some cards over. Having flipped them, you then read the numbers facing up and recalculate the score. What is the maximum score you can achieve?

Write a function:

func Solution(A []int, B []int) int

that, given two arrays of integers A and B, both of length N, describing the numbers written on both sides of the cards, facing up and down respectively, returns the maximum possible score.

For example, given A = [1, 2, 4, 3] and B = [1, 3, 2, 3], your function should return 5, as without flipping any card the smallest positive integer excluded from this sequence is 5.

Given A = [4, 2, 1, 6, 5] and B = [3, 2, 1, 7, 7], your function should return 4, as we could flip the first card so that the numbers facing up are [3, 2, 1, 6, 5] and it is impossible to have both numbers 3 and 4 facing up.

Given A = [2, 3] and B = [2, 3] your function should return 1, as no matter how the cards are flipped, the numbers facing up are [2, 3].

Assume that:

N is an integer within the range [1..100,000];
each element of arrays A, B is an integer within the range [1..100,000,000];
input arrays are of equal size.
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: ([1, 2, 4, 3], [1, 3, 2, 3])	0/10

**/

/*
Results:
Best Result of 68% got!
https://app.codility.com/demo/results/training6FQUBK-WY2/
*/

package main

//package solution

// you can also use imports, for example:
import "fmt"

//import "math"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4

	const MAX = 100000 * 2

	N := len(A)

	// First loop, we count
	map_count := make(map[int]int)
	// map1 := make(map[int]map[int]bool)
	map1 := make(map[int]int)
	//map2 := make(map[int]int)
	for i := 0; i < N; i++ {
		a := A[i]
		b := B[i]

		if a == b {
			map_count[a] = MAX
			//add_to_map(map1, a, a)
			map1[a] = a
			continue
		}

		v, ok := map_count[a]

		if !ok {
			map_count[a] = 1
		} else {
			map_count[a] = v + 1
		}

		v, ok = map_count[b]
		if !ok {
			map_count[b] = 1
		} else {
			map_count[b] = v + 1
		}

		// will be helpful only for count 1 cases e.g 2/3 or 2/7
		//add_to_map(map1, a, b)
		//add_to_map(map1, b, a)
		map1[a] = b
		map1[b] = a
	}

	map_count2 := make(map[int]int)
	for k, v := range map_count {
		if v == 1 {
			map_count2[k] = v
		}
	}

	//fmt.Printf("map_count: %v\n", map_count)
	fmt.Printf("map_count2: %v\n", map_count2)

	// 2nd loop we find the max possible val
	max_val := N + 1
	for i := 1; i <= N; i++ {
		_, ok := map_count[i]
		if !ok {
			max_val = i
			break
		}
	}

	// in the 3rd lopp we find counts == 1 upto max_val
	ret := max_val
	for a := 1; a < ret; a++ {
		if map_count[a] == 0 {
			ret = a
			break
		}

		_, ok := map_count2[a]
		if !ok {
			continue
		}

		b := map1[a]
		map_count[b] -= 1
		if map_count[b] == 0 {
			if b < a {
				ret = a
				break
			}
		}

	}

	return ret
}

func add_to_map(map1 map[int]map[int]bool, key, val int) {
	map2, ok := map1[key]
	if ok {
		map2[val] = true
	} else {
		map2 := make(map[int]bool)
		map2[val] = true
		map1[key] = map2
	}

}

func main() {
	A := []int{1, 1}
	B := []int{2, 3}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret := Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{1, 1}
	B = []int{2, 1}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{1, 2, 3}
	B = []int{2, 1, 1}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{1, 2, 3, 4, 5}
	B = []int{1, 1, 2, 2, 3}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{4, 2, 1, 6, 5}
	B = []int{3, 2, 1, 7, 7}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{1, 2, 3, 4, 1, 7, 7}
	B = []int{1, 5, 5, 5, 1, 2, 3}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

	A = []int{1, 2, 2, 4, 1}
	B = []int{1, 3, 5, 4, 1}
	fmt.Printf("A: %v\nB: %v\n", A, B)
	ret = Solution(A, B)
	fmt.Printf("ret: %v\n", ret)

}
