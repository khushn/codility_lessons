/**
You are given N counters, initially set to 0, and you have two possible operations on them:

increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.
A non-empty array A of M integers is given. This array represents consecutive operations:

if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.
For example, given integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the values of the counters after each consecutive operation will be:

    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)
The goal is to calculate the value of every counter after all operations.

Write a function:

func Solution(N int, A []int) []int

that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

The sequence should be returned as:

a structure Results (in C), or
a vector of integers (in C++), or
a record Results (in Pascal), or
an array of integers (in any other programming language).
For example, given:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the function should return [3, 2, 2, 4, 2], as explained above.

Assume that:

N and M are integers within the range [1..100,000];
each element of array A is an integer within the range [1..N + 1].
Complexity:

expected worst-case time complexity is O(N+M);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: (5, [3, 4, 4, 6, 1, 4, 4]) 0/10
Solution
Solution
Go 1.4

Run Tests

Submit


14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
    row_max := 0
    ret := make([]int, N)
    for _, v := range A {
        if v > N {
            global_max := row_max
        } else {
            if ret[i] == 0 {
                ret[i]= global_max+1
            } else {
                ret[i]++
                if ret[i] > row_max {
                    row_max = ret[i]
                }
            }
        }
    }

    for i:=0; i<N; i++ {
        if ret[i] == 0 {
            ret[i]=global_max
        }
    }
    return ret
}
All changes saved
Test Output
Test Output
**/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4

	global_max := 0
	global_max_pos := -1
	thread_max := 0
	my_map := make(map[int]int)
	for i, v := range A {
		if v <= N {
			count, ok := my_map[v]
			if !ok {
				my_map[v] = 1
				if thread_max < 1 {
					thread_max = 1
				}
			} else {
				tmp := count + 1
				my_map[v] = tmp
				if thread_max < tmp {
					thread_max = tmp
				}
			}
		} else {
			global_max += thread_max
			global_max_pos = i
			// reset
			my_map = make(map[int]int)
			thread_max = 0
		}
	}

	ret := make([]int, N)

	for i := 0; i < N; i++ {
		ret[i] = global_max
	}

	for i := global_max_pos + 1; i < len(A); i++ {
		ret[A[i]-1]++
	}
	return ret
}
