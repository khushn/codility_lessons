/**
A prime is a positive integer X that has exactly two distinct divisors: 1 and X. The first few prime integers are 2, 3, 5, 7, 11 and 13.

A prime D is called a prime divisor of a positive integer P if there exists a positive integer K such that D * K = P. For example, 2 and 5 are prime divisors of 20.

You are given two positive integers N and M. The goal is to check whether the sets of prime divisors of integers N and M are exactly the same.

For example, given:

N = 15 and M = 75, the prime divisors are the same: {3, 5};
N = 10 and M = 30, the prime divisors aren't the same: {2, 5} is not equal to {2, 3, 5};
N = 9 and M = 5, the prime divisors aren't the same: {3} is not equal to {5}.
Write a function:

func Solution(A []int, B []int) int

that, given two non-empty arrays A and B of Z integers, returns the number of positions K for which the prime divisors of A[K] and B[K] are exactly the same.

For example, given:

    A[0] = 15   B[0] = 75
    A[1] = 10   B[1] = 30
    A[2] = 3    B[2] = 5
the function should return 1, because only one pair (15, 75) has the same set of prime divisors.

Assume that:

Z is an integer within the range [1..6,000];
each element of arrays A, B is an integer within the range [1..2,147,483,647].
Complexity:

expected worst-case time complexity is O(Z*log(max(A)+max(B))2);
expected worst-case space complexity is O(1) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: ([15, 10, 9], [75, 30, 5])	0/10

*/

/**
Solution:
Corrent one is only 76% (so not a perfect one)
https://app.codility.com/demo/results/training4KMVMT-UZN/

Update (in the evening on 16th May 2018):
I have now a 100% solution!
https://app.codility.com/demo/results/trainingGBP5JF-BXC/
**/

package main

// you can also use imports, for example:
import "fmt"

// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4

	N := len(A)

	ret := 0

for_label:
	for i := 0; i < N; i++ {
		a := A[i]
		b := B[i]
		gcd1 := gcd(a, b)
		min_ab := int(math.Min(float64(a), float64(b)))
		max_ab := int(math.Max(float64(a), float64(b)))
		if min_ab == 1 && max_ab > min_ab {
			continue
		}

		if gcd1 == 1 && min_ab > 1 {
			continue
		}

		arem := a / gcd1
		for arem > 1 {
			gcd2 := gcd(arem, gcd1)
			if gcd2 == 1 {
				continue for_label
			}
			arem /= gcd2
		}
		brem := b / gcd1
		for brem > 1 {
			gcd2 := gcd(brem, gcd1)
			if gcd2 == 1 {
				continue for_label
			}
			brem /= gcd2
		}
		fmt.Printf("a: %v, b: %v\n", a, b)
		ret++
	}

	return ret
}

// Greatest common divisor
func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		if a%b == 0 {
			return b
		}
		return gcd(a%b, b)
	} else {
		if b%a == 0 {
			return a
		}
		return gcd(a, b%a)
	}
}

func main() {
	var A []int
	var B []int
	for i := 1; i < 71; i++ {
		for j := 1; j < 71; j++ {
			A = append(A, i)
			B = append(B, j)
		}
	}

	fmt.Printf("No. of pairs: %v\n", len(A))

	ret := Solution(A, B)
	fmt.Printf("Result: %v\n", ret)
}
