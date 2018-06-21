/**
On the sequence of logical values (true or false) we can build up an OR-Pascal-triangle structure. Instead of summing the values, as in a standard Pascal-triangle, we will combine them using the OR function. That means that the lowest row is simply the input sequence, and every entry in each subsequent row is the OR of the two elements below it. For example, the OR-Pascal-triangle built on the array [true, false, false, true, false] is as follows:

Your job is to count the number of nodes in the OR-Pascal-triangle that contain the value true (this number is 11 for the animation above).

Write a function:

func Solution(P []bool) int

that, given an array P of N Booleans, returns the number of fields in the OR-Pascal-triangle built on P that contain the value true. If the result is greater than 1,000,000,000, your function should return 1,000,000,000.

Given P = [true, false, false, true, false], the function should return 11, as explained above.

Given P = [true, false, false, true], the function should return 7, as can be seen in the animation below.

Assume that:

N is an integer within the range [1..100,000].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Solution
Programming language used
Go
Total time used
9 minutes
Effective time used
9 minutes
Notes
not defined yet
Task timeline


**/

/**
Result: Golden award
https://app.codility.com/cert/view/certMFD2TT-RF3AM9DFM24TKAYU/details/
**/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(P []bool) int {
	// write your code in Go 1.4
	n := len(P)
	s := n * (n + 1) / 2

	falseS := 0
	falseC := 0
	for i := 0; i < n; i++ {
		if P[i] == false {
			falseC++
		} else {
			if falseC == 1 {
				falseS++
			} else if falseC >= 2 {
				falseS += (falseC + 1) * falseC / 2
			}
			falseC = 0
		}
	}

	if falseC == 1 {
		falseS++
	} else if falseC >= 2 {
		falseS += (falseC + 1) * falseC / 2
	}

	res := s - falseS

	LIM := 1000000000
	if res > LIM {
		res = LIM
	}
	return res
}
