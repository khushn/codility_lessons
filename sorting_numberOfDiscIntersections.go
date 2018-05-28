/**
We draw N discs on a plane. The discs are numbered from 0 to N − 1. An array A of N non-negative integers, specifying the radiuses of the discs, is given. The J-th disc is drawn with its center at (J, 0) and radius A[J].

We say that the J-th disc and K-th disc intersect if J ≠ K and the J-th and K-th discs have at least one common point (assuming that the discs contain their borders).

The figure below shows discs drawn for N = 6 and A as follows:

  A[0] = 1
  A[1] = 5
  A[2] = 2
  A[3] = 1
  A[4] = 4
  A[5] = 0


There are eleven (unordered) pairs of discs that intersect, namely:

discs 1 and 4 intersect, and both intersect with all the other discs;
disc 2 also intersects with discs 0 and 3.
Write a function:

func Solution(A []int) int

that, given an array A describing N discs as explained above, returns the number of (unordered) pairs of intersecting discs. The function should return −1 if the number of intersecting pairs exceeds 10,000,000.

Given array A shown above, the function should return 11, as explained above.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [0..2,147,483,647].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [1, 5, 2, 1, 4, 0]	0/10

*/

package solution

// you can also use imports, for example:
//import "fmt"
import "sort"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type CirclePoint struct {
	X        int
	CircleId int
	IsLeft   bool
}

// ByLeftX implements sort.Interface for []*CirclePoint based on
// the Left x value field.
type ByLeftX []*CirclePoint

func (a ByLeftX) Len() int      { return len(a) }
func (a ByLeftX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLeftX) Less(i, j int) bool {
	// a crucial logic is that if there is a clash of left and right circle
	// then we should return the left one first, as touching counts as intersection
	ret := false
	if a[i].X < a[j].X {
		ret = true
	} else if a[i].X == a[j].X {
		if a[i].IsLeft {
			ret = true
		}
	}
	return ret
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var cparr []*CirclePoint
	N := len(A)
	for i := 0; i < N; i++ {
		cpl := CirclePoint{i - A[i], i, true}
		cpr := CirclePoint{i + A[i], i, false}
		cparr = append(cparr, &cpl)
		cparr = append(cparr, &cpr)
	}

	// sort the left and right points on the cricle starting from left to right
	sort.Sort(ByLeftX(cparr))
	//fmt.Printf("cparr: %+v\n", cparr[0])

	// Now we do some prefix sum to avoid time complexity run to O(N^2)
	// we just keep the sum of left edges of a circle enountered so far
	pref_sum := make([]int, 2*N)
	for i := 0; i < 2*N; i++ {
		if i > 0 {
			pref_sum[i] = pref_sum[i-1]
		}
		cp := cparr[i]
		if cp.IsLeft {
			pref_sum[i]++
		} else {
			pref_sum[i]--
		}
	}

	// Find the intersections of all left sides with the already encountered left sides,
	// as we move from left to right
	// we keep a track of the circles we are in
	ret := 0
	for i := 0; i < 2*N; i++ {
		open_circles := pref_sum[i]
		cp := cparr[i]
		if cp.IsLeft {
			// minus 1 as we need at least one outer circle always open
			ret += open_circles - 1
			if ret > 10000000 {
				return -1
			}
		}
	}

	return ret
}
