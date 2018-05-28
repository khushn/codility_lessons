/**
You have to climb up a ladder. The ladder has exactly N rungs, numbered from 1 to N. With each step, you can ascend by one or two rungs. More precisely:

with your first step you can stand on rung 1 or 2,
if you are on rung K, you can move to rungs K + 1 or K + 2,
finally you have to stand on rung N.
Your task is to count the number of different ways of climbing to the top of the ladder.

For example, given N = 4, you have five different ways of climbing, ascending by:

1, 1, 1 and 1 rung,
1, 1 and 2 rungs,
1, 2 and 1 rung,
2, 1 and 1 rungs, and
2 and 2 rungs.
Given N = 5, you have eight different ways of climbing, ascending by:

1, 1, 1, 1 and 1 rung,
1, 1, 1 and 2 rungs,
1, 1, 2 and 1 rung,
1, 2, 1 and 1 rung,
1, 2 and 2 rungs,
2, 1, 1 and 1 rungs,
2, 1 and 2 rungs, and
2, 2 and 1 rung.
The number of different ways can be very large, so it is sufficient to return the result modulo 2P, for a given integer P.

Write a function:

func Solution(A []int, B []int) []int

that, given two non-empty arrays A and B of L integers, returns an array consisting of L integers specifying the consecutive answers; position I should contain the number of different ways of climbing the ladder with A[I] rungs modulo 2B[I].

For example, given L = 5 and:

    A[0] = 4   B[0] = 3
    A[1] = 4   B[1] = 2
    A[2] = 5   B[2] = 4
    A[3] = 5   B[3] = 3
    A[4] = 1   B[4] = 1
the function should return the sequence [5, 1, 8, 0, 1], as explained above.

Assume that:

L is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..L];
each element of array B is an integer within the range [1..30].
Complexity:

expected worst-case time complexity is O(L);
expected worst-case space complexity is O(L) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
*/

/*
Worked well, in first try
https://app.codility.com/demo/results/trainingHJVSTU-6AC/
*/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) []int {
	// write your code in Go 1.4

	// Fibonacci series can be applied to the above ladder climbing problem
	// if we want to climb n steps then
	// f(n) = f(n-1) + f(n-2)
	// Logic f(n-2) already investigates all the ways upto step n-2, so after that there is
	// only one unique way by taking a 2 step to go onto nth,
	// (note we can do 1+1 afer step n-2, but that permutation would have been covered in f(n-1))
	// From f(n-1), we have only one way to climb and that is by taking a single step

	MAX_NUM := 60000

	max_pow := int(math.Pow(float64(2), float64(30)))

	fib_arr := make([]int, MAX_NUM+1)
	fib_arr[0] = 1
	fib_arr[1] = 1
	for i := 2; i < MAX_NUM+1; i++ {
		// We take the remainder from the largest number
		// since we can again take it with a lesser power of 2, if needed
		// e.g. if we take remainder/modulo from 2^3 i.e. 8, we can again
		// safelt take it from 2^2 i.e. 4, and get the same answer as if had not
		// taken it from 8 earlier.
		fib_arr[i] = (fib_arr[i-1] + fib_arr[i-2]) % max_pow
	}

	N := len(A)
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		a := A[i]
		b := B[i]
		pow_b := int(math.Pow(float64(2), float64(b)))
		ret[i] = fib_arr[a] % pow_b
	}

	return ret
}