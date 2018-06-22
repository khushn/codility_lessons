/**
Codility's Cobaltium 2018 challenge

You have two sequences A and B consisting of integers, both of length N, and you would like them to be (strictly) increasing, i.e. for each K (0 ≤ K < N − 1), A[K] < A[K + 1] and B[K] < B[K + 1]. Thus, you need to modify the sequences, but the only manipulation you can perform is to swap an arbitrary element in sequence A with the corresponding element in sequence B. That is, both elements to be exchanged must occupy the same index position within each sequence.

For example, given A = [5, 3, 7, 7, 10] and B = [1, 6, 6, 9, 9], you can swap elements at positions 1 and 3, obtaining A = [5, 6, 7, 9, 10], B = [1, 3, 6, 7, 9].

Your goal is make both sequences increasing, using the smallest number of moves.

Write a function:

func Solution(A []int, B []int) int

that, given two arrays A, B of length N, containing integers, returns the minimum number of swapping operations required to make the given arrays increasing. If it is impossible to achieve the goal, return −1.

For example, given:

A[0] = 5        B[0] = 1
A[1] = 3        B[1] = 6
A[2] = 7        B[2] = 6
A[3] = 7        B[3] = 9
A[4] = 10       B[4] = 9
your function should return 2, as explained above.

Given:

A[0] = 5        B[0] = 2
A[1] = -3       B[1] = 6
A[2] = 6        B[2] = -5
A[3] = 4        B[3] = 1
A[4] = 8        B[4] = 0
your function should return −1, since you cannot perform operations that would make the sequences become increasing.

Given:

A[0] = 1        B[0] = -2
A[1] = 5        B[1] = 0
A[2] = 6        B[2] = 2
your function should return 0, since the sequences are already increasing.

Assume that:

N is an integer within the range [2..100,000];
each element of arrays A, B is an integer within the range [−1,000,000,000..1,000,000,000];
A and B have equal lengths.
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: ([5, 3, 7, 7, 10], [1, 6, 6, 9, 9])	0/10

**/

/**

Result: This one gets 100%. Happy!

https://app.codility.com/demo/results/training9ZCW25-B65/
*/

package main

// package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4

	N := len(A)

	for i := 0; i < N-1; i++ {
		if A[i]+B[i] >= A[i+1]+B[i+1] {
			return -1
		}
	}

	// get all the good positions i.e.
	// where both A[i], B[i] are less than A[i+1] & B[i+1]
	// these are like good/safe barriers, which contain the spread of
	// change
	var good []int
	for i := 0; i < N-1; i++ {
		if A[i] < A[i+1] && A[i] < B[i+1] && B[i] < B[i+1] && B[i] < A[i+1] {
			good = append(good, i+1)
		}

	}

	good = append(good, N)
	ngood := len(good)
	//fmt.Printf("ngood: %v\n", ngood)
	ret := 0
	start := 0
	for j := 0; j < ngood; j++ {
		end := good[j]
		tret := 0
		for i := start + 1; i < end; i++ {
			if A[i] <= A[i-1] || B[i] <= B[i-1] {
				tmp := A[i]
				A[i] = B[i]
				B[i] = tmp
				tret++
			}
		}
		//fmt.Printf("tret: %v\n", tret)
		dist := end - start
		if tret > dist/2 {
			tret = dist - tret
		}
		start = end
		ret += tret
	}

	return ret
}

func main() {
	var arr1 = []int{3, 5}
	var arr2 = []int{4, 4}
	ret := Solution(arr1, arr2)
	fmt.Printf("ret: %v\n", ret)
}
