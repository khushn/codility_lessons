/**

Write a function:

func Solution(A []int) int

that, given an array of N positive integers, returns the maximum number of trailing zeros of the number obtained by multiplying three different elements from the array. Numbers are considered different if they are at different positions in the array.

For example, given A = [7, 15, 6, 20, 5, 10], the function should return 3 (you can obtain three trailing zeros by taking the product of numbers 15, 20 and 10 or 20, 5 and 10).

For another example, given A = [25, 10, 25, 10, 32], the function should return 4 (you can obtain four trailing zeros by taking the product of numbers 25, 25 and 32).

Assume that:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [1..1,000,000,000].
Complexity:

expected worst-case time complexity is O(N*log(max(A)));
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [7, 15, 6, 20, 5, 10]	0/10

**/

/**
Acknowldegement: My son, helped me with this. And it was great to crack this together
Result: Golden award
https://app.codility.com/cert/view/cert2WGNS7-ATHDKRBKHASKYRYU/
*/

package main

//package solution

// you can also use imports, for example:
import "fmt"

// import "os"
// import "sort"

import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// These limits are based on
// 5^14 is around 1 billion (max value can be 1 billion)
const MAX5 = 14

// 27 is the max no. of zeros if all 3 numbers are 1 billion
const MAX2 = 28

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)

	dp := make([]int, MAX5*MAX2)

	for i := 0; i < N; i++ {
		a := A[i]

		ftmp := int(math.Min(float64(MAX5-1), float64(num_divs(a, 5))))
		//dp5[a] = ftmp
		ttmp := int(math.Min(float64(MAX2-1), float64(num_divs(a, 2))))
		//dp2[a] = ttmp
		dp[ftmp*MAX2+ttmp]++
	}

	// fmt.Printf("dp: %v\n", dp)

	ret := 0
	for i := 0; i < MAX5*MAX2-2; i++ {
		if dp[i] == 0 {
			continue
		}
		for j := i + 1; j < MAX5*MAX2-1; j++ {
			if dp[j] == 0 {
				continue
			}
			for k := j + 1; k < MAX5*MAX2; k++ {
				if dp[k] == 0 {
					continue
				}
				fives := i/MAX2 + j/MAX2 + k/MAX2
				twos := i%MAX2 + j%MAX2 + k%MAX2
				tmp := int(math.Min(float64(fives), float64(twos)))
				if tmp > ret {
					ret = tmp
				}
			}
		}
	}

	// consider 3 or more at same position
	for i := 0; i < MAX5*MAX2; i++ {
		if dp[i] >= 3 {
			fives := 3 * (i / MAX2)
			twos := 3 * (i % MAX2)
			tmp := int(math.Min(float64(fives), float64(twos)))
			if tmp > ret {
				ret = tmp
			}
		}
	}

	// now considering only 2 at a given slot
	for i := 0; i < MAX5*MAX2; i++ {
		if dp[i] >= 2 {
			for j := 0; j < MAX5*MAX2; j++ {
				if j != i {
					if dp[j] == 0 {
						continue
					}

					fives := 2*(i/MAX2) + j/MAX2
					twos := 2*(i%MAX2) + j%MAX2
					tmp := int(math.Min(float64(fives), float64(twos)))
					if tmp > ret {
						ret = tmp
					}
				}

			}
		}
	}

	return ret

}

func num_divs(a, d int) int {
	ret := 0
	for a%d == 0 && a > 1 {
		a = a / d
		ret++
	}
	return ret
}

func main() {
	var test = []int{125, 100, 125, 100, 64}
	fmt.Printf("test: %v\n", test)
	res := Solution(test)
	fmt.Printf("res: %v\n", res)

	var test2 = []int{int(math.Pow(5, 13)), int(math.Pow(2, 10)), 8, 10000}
	fmt.Printf("test2: %v\n", test2)
	res = Solution(test2)
	fmt.Printf("res: %v\n", res)

	var test3 = []int{1, 1, 101}
	fmt.Printf("test3: %v\n", test3)
	res = Solution(test3)
	fmt.Printf("res: %v\n", res)

	var test4 = []int{5, 5, 4}
	fmt.Printf("test4: %v\n", test4)
	res = Solution(test4)
	fmt.Printf("res: %v\n", res)

	var test5 = []int{8, 1, 1}
	fmt.Printf("test5: %v\n", test5)
	res = Solution(test5)
	fmt.Printf("res: %v\n", res)
}
