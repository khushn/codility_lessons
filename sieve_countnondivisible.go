/*
You are given an array A consisting of N integers.

For each number A[i] such that 0 ≤ i < N, we want to count the number of elements of the array that are not the divisors of A[i]. We say that these elements are non-divisors.

For example, consider integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
For the following elements:

A[0] = 3, the non-divisors are: 2, 6,
A[1] = 1, the non-divisors are: 3, 2, 3, 6,
A[2] = 2, the non-divisors are: 3, 3, 6,
A[3] = 3, the non-divisors are: 2, 6,
A[4] = 6, there aren't any non-divisors.
Write a function:

func Solution(A []int) []int

that, given an array A consisting of N integers, returns a sequence of integers representing the amount of non-divisors.

The sequence should be returned as:

a structure Results (in C), or
a vector of integers (in C++), or
a record Results (in Pascal), or
an array of integers (in any other programming language).
For example, given:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
the function should return [2, 4, 3, 2, 0], as explained above.

Assume that:

N is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..2 * N].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [3, 1, 2, 3, 6]	0/10

*/

/**

RESULT: Got perfect score for this one!
https://app.codility.com/demo/results/training8HMSNY-JYV/
**/

package solution

// you can also use imports, for example:
//import "fmt"

// import "os"
//import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) []int {
	// write your code in Go 1.4

	N := len(A)

	mymap := make(map[int]int)
	for i := 0; i < N; i++ {
		_, ok := mymap[A[i]]
		if !ok {
			mymap[A[i]] = 1
		} else {
			mymap[A[i]] += 1
		}
	}

	// construct a custom Eratosthenes sieve
	// of factors
	// Note we create a custom sieve based on numbers we have
	// in the map created above
	sieve := make([]int, 2*N+1)
	for i := 1; i <= 2*N; i++ {
		v, ok := mymap[i]
		if ok {
			for k := i; k <= 2*N; k += i {
				sieve[k] += v
			}
		}
	}

	// Now we have every thing computed in the sieve,
	// with the help of the sieve.
	// We just loop again,
	// and return the values from the sive

	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = N - sieve[A[i]]
	}

	return ret

}
