/**
The Fibonacci sequence is defined using the following recursive formula:

    F(0) = 0
    F(1) = 1
    F(M) = F(M - 1) + F(M - 2) if M >= 2
A small frog wants to get to the other side of a river. The frog is initially located at one bank of the river (position −1) and wants to get to the other bank (position N). The frog can jump over any distance F(K), where F(K) is the K-th Fibonacci number. Luckily, there are many leaves on the river, and the frog can jump between the leaves, but only in the direction of the bank at position N.

The leaves on the river are represented in an array A consisting of N integers. Consecutive elements of array A represent consecutive positions from 0 to N − 1 on the river. Array A contains only 0s and/or 1s:

0 represents a position without a leaf;
1 represents a position containing a leaf.
The goal is to count the minimum number of jumps in which the frog can get to the other side of the river (from position −1 to position N). The frog can jump between positions −1 and N (the banks of the river) and every position containing a leaf.

For example, consider array A such that:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
The frog can make three jumps of length F(5) = 5, F(3) = 2 and F(5) = 5.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the minimum number of jumps by which the frog can get to the other side of the river. If the frog cannot reach the other side of the river, the function should return −1.

For example, given:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
the function should return 3, as explained above.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0]	0/10

*/

/*
Results: 91% ( I guess, due to use of recursion. Its not reaching 100%)
https://app.codility.com/demo/results/training8SRH96-URZ/
*/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

var dp_map map[int]int

func Solution(A []int) int {
	// write your code in Go 1.4

	dp_map = make(map[int]int)

	// We cache fibonacci numbers upto max val i.e. 100,000
	MAX_VAL := 100001
	var fib_arr []int
	fib_arr = append(fib_arr, 0)
	fib_arr = append(fib_arr, 1)
	for i := 2; fib_arr[i-1] < MAX_VAL; i++ {
		tmp := fib_arr[i-1] + fib_arr[i-2]
		fib_arr = append(fib_arr, tmp)
	}

	N := len(A)
	// find leave positions
	first_leave := -1
	leave_map := make(map[int]bool)
	for i, v := range A {
		if v == 1 {
			if first_leave == -1 {
				first_leave = i
			}
			leave_map[i] = true
		}
	}

	return jump_possible(-1, first_leave, N, leave_map, fib_arr)
}

func jump_possible(cur_pos, first_leave, N int, leave_map map[int]bool, fib_arr []int) int {
	X := len(fib_arr)

	min_jumps := -1

	for i := X - 1; i >= 1; i-- {
		dist := N - cur_pos
		if fib_arr[i] == dist {
			// we can go to the destination in a single jump from here
			return 1
		} else if fib_arr[i] < dist {
			test_pos := cur_pos + fib_arr[i]
			if test_pos < first_leave {
				break
			}
			if leave_map[test_pos] {
				// possible jump upto here
				// but check recursively if jump possible to dest from here
				mid_ret, ok := dp_map[test_pos]
				if !ok {
					mid_ret = jump_possible(test_pos, first_leave, N, leave_map, fib_arr)
					dp_map[test_pos] = mid_ret
				}
				if mid_ret > 0 {
					// we need one jump to here
					// and add the mid_ret to it i.e. no. of steps to destination
					if min_jumps == -1 || mid_ret+1 < min_jumps {
						min_jumps = mid_ret + 1
					}
				}

			}
		}
	}
	return min_jumps
}
