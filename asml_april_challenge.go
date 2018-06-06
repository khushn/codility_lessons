/**
A long time ago, when the most basic model of an HP3000 computer used to cost $95,000 (over half a million in today's dollars), a very wise man called Gordon Moore made predictions about how computers would become cheaper and more powerful each year. According to Moore, the number of components per integrated circuit would double every two years. Thanks to the creative, determined engineers working in circuit printing technology, we do indeed have smaller, cheaper and more powerful computers today.

Circuit printing, as we call photolithography, is an extremely complex yet rewarding field, and ASML needs the best software engineers in the world to make this magic happen. We work closely with our clients to help them print their circuits in the most effective way. One of our clients requests us to write a method to optimize their circuit efficiency. The circuit is represented as a string consisting of the letters "M" and "L", where M represents Memory units and L represents Logic units. The efficiency of the circuit is measured as the length of the longest interval of letters "M". For example, given input string "LMMMLMMMMLLLM", the longest interval is 4.

Our customer wants to change the circuit in such a way that the longest M-interval will be equal to K. We can change any unit at any position in the circuit, i.e. either we can change any "M" to "L" or any "L" to "M". The objective of this challenge is to calculate the minimum number of changes we have to make in order to achieve the desired longest M-interval length K.

Write a function:

func Solution(S string, K int) int

where the first argument, S, represents the circuit as a string of length N that consists of only characters "M" and/or "L" and the second argument, K, is the desired longest M-interval in the string. The return value shall be the minimum number of changes to achieve K as the longest M-interval in the input string.

For example, given S = "MLMMLLM" and K = 3, your function should return 1. We can change the letter at position 4 (counting from 0) to obtain "MLMMMLM", in which the longest interval of letters "M" is exactly three characters long.

For another example, given S = "MLMMMLMMMM" and K = 2, your function should return 2. We can, for example, modify the letters at positions 2 and 7 to get the string "MLLMMLMLMM", which satisfies the desired property.

Assume that:

string S consists only of the characters "M" and/or "L";
N is an integer within the range [1..100,000];
K is an integer within the range [0..N].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

**/

/**
Results:
Got 100% after a few tries
https://app.codility.com/demo/results/trainingEJ4MFB-BC7/
**/

package main

//package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, K int) int {
	// write your code in Go 1.4

	N := len(S)

	// boundary case 1
	if K == 0 {
		ret := 0
		for i := 0; i < N; i++ {
			ch := S[i]
			if ch == 'M' {
				ret++
			}
		}
		return ret
	}

	arr1 := make([]int, N)

	maxl := 0
	tc := -1
	prevm := false
	for i := 0; i < N; i++ {
		if S[i] == 'M' {
			if prevm {
				tc++
			} else {
				tc = 1
			}
			if tc > maxl {
				maxl = tc
			}
			arr1[i] = tc
			prevm = true

		} else {
			arr1[i] = 0
			prevm = false
		}

	}

	// calculate M count, for K sliding window
	arr2 := make([]int, N)
	init_sum := 0
	for i := 0; i < K; i++ {
		if S[i] == 'M' {
			init_sum++
		}
		arr2[i] = init_sum
	}

	for i := K; i < N; i++ {
		arr2[i] = arr2[i-1]
		if S[i-K] == 'M' {
			arr2[i]--
		}

		if S[i] == 'M' {
			arr2[i]++
		}
	}

	fmt.Println("arr2: %v\n", arr2)

	if K == maxl {
		return 0
	} else if K < maxl {
		ret := 0
		for i := 1; i < N; i++ {
			if arr1[i] == 0 && arr1[i-1] > 0 {
				if arr1[i-1] > K {
					ret += arr1[i-1] / (K + 1)
				}
			}
		}

		if arr1[N-1] > K {
			ret += arr1[N-1] / (K + 1)
		}
		return ret
	} else {

		minret := N

		for i := K - 1; i < N; i++ {
			tmpret := K - arr2[i]
			if i < N-1 && S[i+1] == 'M' {
				tmpret++
			}

			if i-K >= 0 && S[i-K] == 'M' {
				tmpret++
			}

			if tmpret < minret {
				minret = tmpret
			}

		}
		return minret
	}

	return N
}

func main() {
	// ('LLLMMLLLLLLMMMLLLLL', 5)
	S := "LLLMMLLLLLLMMMLLLLL"
	ret := Solution(S, 5)
	fmt.Printf("S: %v, ret: %v\n", S, ret)

	S = "MLM"
	ret = Solution(S, 2)
	fmt.Printf("S: %v, ret: %v\n", S, ret)

}
