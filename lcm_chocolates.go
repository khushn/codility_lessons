/**
Two positive integers N and M are given. Integer N represents the number of chocolates arranged in a circle, numbered from 0 to N − 1.

You start to eat the chocolates. After eating a chocolate you leave only a wrapper.

You begin with eating chocolate number 0. Then you omit the next M − 1 chocolates or wrappers on the circle, and eat the following one.

More precisely, if you ate chocolate number X, then you will next eat the chocolate with number (X + M) modulo N (remainder of division).

You stop eating when you encounter an empty wrapper.

For example, given integers N = 10 and M = 4. You will eat the following chocolates: 0, 4, 8, 2, 6.

The goal is to count the number of chocolates that you will eat, following the above rules.

Write a function:

func Solution(N int, M int) int

that, given two positive integers N and M, returns the number of chocolates that you will eat.

For example, given integers N = 10 and M = 4. the function should return 5, as explained above.

Assume that:

N and M are integers within the range [1..1,000,000,000].
Complexity:

expected worst-case time complexity is O(log(N+M));
expected worst-case space complexity is O(log(N+M)).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: (10, 4)	0/10

*/

/**

Results here:
https://app.codility.com/demo/results/trainingF73QS2-MQ5/
(This problem was in PAINLESS category, which is supposed to be easy)

**/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, M int) int {
	// write your code in Go 1.4

	// We need to calculate the LCM of the two no.s and divide it by M

	// LCM is equal to N*M/gcd(N. N)
	lcm := N * M / gcd(N, M)
	return lcm / M
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
