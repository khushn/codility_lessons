/*
You are given two non-empty arrays A and B consisting of N integers. These arrays represent N planks. More precisely, A[K] is the start and B[K] the end of the K−th plank.

Next, you are given a non-empty array C consisting of M integers. This array represents M nails. More precisely, C[I] is the position where you can hammer in the I−th nail.

We say that a plank (A[K], B[K]) is nailed if there exists a nail C[I] such that A[K] ≤ C[I] ≤ B[K].

The goal is to find the minimum number of nails that must be used until all the planks are nailed. In other words, you should find a value J such that all planks will be nailed after using only the first J nails. More precisely, for every plank (A[K], B[K]) such that 0 ≤ K < N, there should exist a nail C[I] such that I < J and A[K] ≤ C[I] ≤ B[K].

For example, given arrays A, B such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10
four planks are represented: [1, 4], [4, 5], [5, 9] and [8, 10].

Given array C such that:

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
if we use the following nails:

0, then planks [1, 4] and [4, 5] will both be nailed.
0, 1, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, 3, then all the planks will be nailed.
Thus, four is the minimum number of nails that, used sequentially, allow all the planks to be nailed.

Write a function:

func Solution(A []int, B []int, C []int) int

that, given two non-empty arrays A and B consisting of N integers and a non-empty array C consisting of M integers, returns the minimum number of nails that, used sequentially, allow all the planks to be nailed.

If it is not possible to nail all the planks, the function should return −1.

For example, given arrays A, B, C such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
the function should return 4, as explained above.

Assume that:

N and M are integers within the range [1..30,000];
each element of arrays A, B, C is an integer within the range [1..2*M];
A[K] ≤ B[K].
Complexity:

expected worst-case time complexity is O((N+M)*log(M));
expected worst-case space complexity is O(M) (not counting the storage required for input arguments).
*/
package solution

// you can also use imports, for example:
//import "fmt"

// import "os"

import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type CC struct {
	V, I int
}

// ByVal implements sort.Interface for []CC based on
// the V value field.
type ByVal []*CC

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i].V < a[j].V }

func Solution(A []int, B []int, C []int) int {
	// write your code in Go 1.4

	N := len(A)
	M := len(C)

	var cc_arr []*CC
	for i, v := range C {
		cc := CC{v, i}
		cc_arr = append(cc_arr, &cc)
	}

	sort.Sort(ByVal(cc_arr))
	//fmt.Printf("cc_arr: %v\n", cc_arr[0])
	max_ind := -1

	for i := 0; i < N; i++ {
		// binary search on cc_arr
		found_start := -1
		//found_end := -1

		// first search for the frame beginning side boundary
		beg := 0
		end := M - 1
		for beg <= end {
			mid := (beg + end) / 2
			v := cc_arr[mid].V
			if v < A[i] {
				beg = mid + 1
			} else {
				// potential beginning found
				found_start = mid
				end = mid - 1
			}
		}

		//fmt.Printf("found_start: %v\n", found_start)
		if found_start == -1 {
			return -1
		}

		// In an earlier iteration, we applied the same technique for pland end pos
		// but then it may be futile,
		// any way we need the minimum index in this frame
		// so we move from here and do it

		if cc_arr[found_start].V > B[i] {
			// if end frame is lesser than this value return early
			return -1
		}

		// else we need to find the best position with least index
		min_tmp_ind := cc_arr[found_start].I
		for j := found_start; j < M && cc_arr[j].V <= B[i]; j++ {
			ind := cc_arr[j].I
			if min_tmp_ind > ind {
				min_tmp_ind = ind
			}
		}

		//fmt.Printf("min_tmp_ind: %v\n", min_tmp_ind)

		if min_tmp_ind > max_ind {
			max_ind = min_tmp_ind
		}
	}

	return max_ind + 1

}
