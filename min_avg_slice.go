/*
*
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
contains the following example slices:

slice (1, 2), whose average is (2 + 2) / 2 = 2;
slice (3, 4), whose average is (5 + 1) / 2 = 3;
slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
The goal is to find the starting position of a slice whose average is minimal.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
the function should return 1, as explained above.

Assume that:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−10,000..10,000].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

*/

/**
Solution notes:
Had to take help of this blog, for the solution:

https://www.martinkysel.com/codility-minavgtwoslice-solution/

Quote from the blog post:
Every slice must be of size two or three. Slices of bigger sizes are created
from such smaller slices. Therefore should any bigger slice have an optimal value,
 all sub-slices must be the same, for this case to hold true. Should this not be
 true, one of the sub-slices must be the optimal slice. The others being bigger.
 Therefore we check all possible slices of size 2/3 and return the smallest one.
 The first such slice is the correct one, do not use <=!

My thoughts:
The case of a sequence of 2 is obvious.
3 is also an atomic one, because of cases like below:
55551515555
In the above 151, offers the minimum average
**/
package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	min_avg := 10001.0
	min_ind := N - 1
	for i := 0; i <= N-2; i++ {
		avgof2 := float64(A[i]+A[i+1]) / 2.0
		if min_avg > avgof2 {
			min_avg = avgof2
			min_ind = i
		}

		if i <= N-3 {
			avgof3 := float64(A[i]+A[i+1]+A[i+2]) / 3.0
			if min_avg > avgof3 {
				min_avg = avgof3
				min_ind = i
			}
		}
	}
	return min_ind
}
