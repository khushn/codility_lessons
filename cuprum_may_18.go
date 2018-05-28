/**
John takes computer security very seriously, so he has picked a very long password to secure his files. However, although the password is very strong, it is also hard to memorize.

John decides to create a new password which is easier to remember. Therefore, it must fulfill certain requirements: he wants each character (digit or letter) to appear in the new password an even number of times (possibly zero). Also, since he is so proud of his original password, he wants the new password to be a contiguous substring of the original password.

John has trouble finding such a substring. Help him by finding, for a given string, the length of the longest substring that fulfills the requirements set out above.

Write a function:

func Solution(S string) int

that, given a non-empty string S consisting of N characters, returns the length of the longest contiguous substring (possibly empty) of string S in which every character occurs an even number of times.

For example, given S = "6daa6ccaaa6d", the function should return 8. The longest valid password taken from S is "a6ccaaa6"; it contains four letters a, and two each of the digit 6 and letter c. Any longer substring must contain either five letters a or one letter d. Given S = "abca", the function should return 0 (note that aa is not a contiguous substring of S).

Assume that:

the length of S is within the range [1..100,000];
S consists only of lowercase letters and digits.
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: '6daa6ccaaa6d'	0/10

**/

/**
Results -- this one got the silver award:

https://app.codility.com/cert/view/cert3MFM5G-ZF3M8NDUY5EVRVH5/


*/
package solution

// you can also use imports, for example:
//import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4

	N := len(S)

	arr := make([][]int, N+1)
	arr[0] = make([]int, 36) // first one all 0's
	for i := 0; i < N; i++ {
		arr[i+1] = make([]int, 36)
		chind := char_val(S[i])
		//fmt.Printf("%v: %v\n", S[i], chind)
		for j := 0; j < 36; j++ {
			arr[i+1][j] = arr[i][j]
		}
		arr[i+1][chind] += 1
	}
	//fmt.Printf("arr0: %v\n", arr[0])
	//fmt.Printf("arrn: %v\n", arr[N])

	ret := 0

	for i := N - 1; i >= 0; i-- {
		for j := -1; j < i && i-j > ret; j++ {
			even := is_even(arr, j+1, i+1)
			if even {
				//fmt.Printf("%v for %v, %v\n", even, i, j)
				if i-j > ret {
					ret = i - j
				}
			}
		}
	}

	return ret
}

func char_val(a byte) int {
	if a >= 'a' && a <= 'z' {
		return int(a - 'a')
	} else {
		return int(26 + a - '0')
	}
}

func is_even(arr [][]int, i, j int) bool {
	arrj := arr[j]
	arri := arr[i]
	for x := 0; x < 36; x++ {
		if arrj[x] > 0 {
			diff := arrj[x] - arri[x]
			if diff%2 != 0 {
				return false
			}
		}
	}

	return true
}
