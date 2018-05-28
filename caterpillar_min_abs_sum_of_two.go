/**
Let A be a non-empty array consisting of N integers.

The abs sum of two for a pair of indices (P, Q) is the absolute value |A[P] + A[Q]|, for 0 ≤ P ≤ Q < N.

For example, the following array A:

  A[0] =  1
  A[1] =  4
  A[2] = -3
has pairs of indices (0, 0), (0, 1), (0, 2), (1, 1), (1, 2), (2, 2).
The abs sum of two for the pair (0, 0) is A[0] + A[0] = |1 + 1| = 2.
The abs sum of two for the pair (0, 1) is A[0] + A[1] = |1 + 4| = 5.
The abs sum of two for the pair (0, 2) is A[0] + A[2] = |1 + (−3)| = 2.
The abs sum of two for the pair (1, 1) is A[1] + A[1] = |4 + 4| = 8.
The abs sum of two for the pair (1, 2) is A[1] + A[2] = |4 + (−3)| = 1.
The abs sum of two for the pair (2, 2) is A[2] + A[2] = |(−3) + (−3)| = 6.
Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the minimal abs sum of two for any pair of indices in this array.

For example, given the following array A:

  A[0] =  1
  A[1] =  4
  A[2] = -3
the function should return 1, as explained above.

Given array A:

  A[0] = -8
  A[1] =  4
  A[2] =  5
  A[3] =-10
  A[4] =  3
the function should return |(−8) + 5| = 3.

Assume that:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [1, 4, -3]	0/10

*/

package main

//package solution

// you can also use imports, for example:
import "fmt"

// import "os"

import "math"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	//N := len(A)

	var neglist []int
	var poslist []int
	for _, v := range A {
		if v == 0 {
			return 0
		} else if v < 0 {
			neglist = append(neglist, -v)
		} else {
			poslist = append(poslist, v)
		}

	}
	sort.Ints(poslist)
	sort.Ints(neglist)

	min := 2000000001

	// check in negative list
	if len(neglist) > 0 && 2*neglist[0] < min {
		min = 2 * neglist[0]
	}

	// check in positive list
	if len(poslist) > 0 && 2*poslist[0] < min {
		min = 2 * poslist[0]
	}

	//negind := 0
	i := 0
	j := 0
	for i < len(poslist) && j < len(neglist) {

		for ; i < len(neglist) && j < len(poslist) && neglist[i] <= poslist[j]; i++ {
			continue
		}

		for ; i < len(neglist) && j < len(poslist) && poslist[j] <= neglist[i]; j++ {

			continue
		}

		if j > 0 && i < len(neglist) && int(math.Abs(float64(neglist[i]-poslist[j-1]))) < min {
			min = int(math.Abs(float64(neglist[i] - poslist[j-1])))
		}

		if i > 0 && j < len(poslist) && int(math.Abs(float64(neglist[i-1]-poslist[j]))) < min {
			min = int(math.Abs(float64(neglist[i-1] - poslist[j])))
		}

		if i < len(neglist) && j < len(poslist) && int(math.Abs(float64(neglist[i]-poslist[j]))) < min {
			min = int(math.Abs(float64(neglist[i] - poslist[j])))
		}

	}

	return min
}

func main() {
	var a = []int{-1000000000, 1000000000}
	ret := Solution(a)
	fmt.Printf("ret: %v\n", ret)
}
