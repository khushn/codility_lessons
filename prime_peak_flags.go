/**
A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbours. More precisely, it is an index P such that 0 < P < N − 1 and A[P − 1] < A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains whose relative heights are represented by array A, as shown in a figure below. You have to choose how many flags you should take with you. The goal is to set the maximum number of flags on the peaks, according to certain rules.



Flags can only be set on peaks. What's more, if you take K flags, then the distance between any two flags should be greater than or equal to K. The distance between indices P and Q is the absolute value |P − Q|.

For example, given the mountain range represented by array A, above, with N = 12, if you take:

two flags, you can set them on peaks 1 and 5;
three flags, you can set them on peaks 1, 5 and 10;
four flags, you can set only three flags, on peaks 1, 5 and 10.
You can therefore set a maximum of three flags in this case.

Write a function:

func Solution(A []int) int

that, given a non-empty array A of N integers, returns the maximum number of flags that can be set on the peaks of the array.

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Assume that:

N is an integer within the range [1..400,000];
each element of array A is an integer within the range [0..1,000,000,000].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2]	0/10

*/

package solution

// you can also use imports, for example:
//import "fmt"

// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	var peaks []int
	num_peaks := 0
	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
			num_peaks++
		}
	}

	if num_peaks == 0 {
		return 0
	}

	dist := peaks[num_peaks-1] - peaks[0]
	sq := math.Sqrt(float64(dist)) + 1
	//fmt.Printf("num_peaks: %v\n", num_peaks)
	//fmt.Printf("sq: %v\n", sq)
	cand := int(math.Min(sq, float64(num_peaks)))

	sq2 := int(math.Sqrt(float64(cand)))

	//max_ret := cand
	for i := cand; i >= sq2; i-- {
		num_used := 1
		prev_peak := peaks[0]
		for j := 1; j < num_peaks && num_used < i; j++ {
			if peaks[j]-prev_peak >= i {
				num_used++
				prev_peak = peaks[j]
			}
		}
		if num_used >= i {
			return i
		}
	}

	return 0
}
