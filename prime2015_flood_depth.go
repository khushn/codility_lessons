/**

You are helping a geologist friend investigate an area with mountain lakes. A recent heavy rainfall has flooded these lakes and their water levels have reached the highest possible point. Your friend is interested to know the maximum depth in the deepest part of these lakes.

We simplify the problem in 2-D dimensions. The whole landscape can be divided into small blocks and described by an array A of length N. Each element of A is the altitude of the rock floor of a block (i.e. the height of this block when there is no water at all). After the rainfall, all the low-lying areas (i.e. blocks that have higher blocks on both sides) are holding as much water as possible. You would like to know the maximum depth of water after this entire area is flooded. You can assume that the altitude outside this area is zero and the outside area can accommodate infinite amount of water.

For example, consider array A such that:

    A[0] = 1
    A[1] = 3
    A[2] = 2
    A[3] = 1
    A[4] = 2
    A[5] = 1
    A[6] = 5
    A[7] = 3
    A[8] = 3
    A[9] = 4
    A[10] = 2
The following picture illustrates the landscape after it has flooded:



The gray area is the rock floor described by the array A above and the blue area with dashed lines represents the water filling the low-lying areas with maximum possible volume. Thus, blocks 3 and 5 have a water depth of 2 while blocks 2, 4, 7 and 8 have a water depth of 1. Therefore, the maximum water depth of this area is 2.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the maximum depth of water.

Given array A shown above, the function should return 2, as explained above.

For the following array:

    A[0] = 5
    A[1] = 8
the function should return 0, because this landscape cannot hold any water.

Assume that:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..100,000,000].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2]  0/10

**/

/**
Results:
https://app.codility.com/demo/results/training3ENP23-EQC/
**/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)

	// max from left
	maxl := make([]int, N)
	for i := 0; i < N; i++ {
		if i == 0 {
			maxl[i] = A[i]
			continue
		}

		if A[i] > maxl[i-1] {
			maxl[i] = A[i]
		} else {
			maxl[i] = maxl[i-1]
		}
	}

	// max from right
	maxr := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		if i == N-1 {
			maxr[i] = A[i]
			continue
		}

		if A[i] > maxr[i+1] {
			maxr[i] = A[i]
		} else {
			maxr[i] = maxr[i+1]
		}
	}

	max_depth := 0
	for i := 0; i < N; i++ {
		bmin := int(math.Min(float64(maxl[i]), float64(maxr[i])))
		depth := bmin - A[i]
		if depth > max_depth {
			max_depth = depth
		}
	}

	return max_depth
}
