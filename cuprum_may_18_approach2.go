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

Solution:
https://app.codility.com/cert/view/certASX472-4AHR5SHDRVQR37YZ/details/

Performance is still O(N^2)

But Silver certificate there, nevertheless

**/
package main

// package solution

// you can also use imports, for example:
import "fmt"

import "math"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type AB struct {
	S int
	E int
}

//var dp_map map[AB]int

var max_len int

func Solution(S string) int {
	// write your code in Go 1.4

	N := len(S)
	max_len = 0

	//dp_map = make(map[AB]int)
	arr := make([][]int, N+1)
	arr[0] = make([]int, 36) // first one all 0'
	arr_pos := make([][]int, 36)
	for i := 0; i < N; i++ {
		arr[i+1] = make([]int, 36)
		chind := char_val(S[i])
		//fmt.Printf("%v: %v\n", S[i], chind)
		for j := 0; j < 36; j++ {
			arr[i+1][j] = arr[i][j]
		}
		arr[i+1][chind] += 1
		arr_pos[chind] = append(arr_pos[chind], i)
		//count_arr[chind] += 1
	}
	//fmt.Printf("arr0: %v\n", arr[0])
	//fmt.Printf("arrn: %v\n", arr[N])

	return recurse_fn(S, 0, N-1, arr, arr_pos)
}

func recurse_fn(S string, s, e int, arr, arr_pos [][]int) int {

	//ab := &AB{s, e}
	//fmt.Printf("looking in db_map: %v for s: %v, e: %v\n", dp_map, s, e)
	//vv, ok2 := dp_map[AB{15, 20}]
	//fmt.Printf("checking retrieval from map: %v, %v\n", ok2, vv)
	//cache_v, ok := dp_map[*ab]

	//if ok {
	//	fmt.Printf("cache hit\n")
	//	return cache_v
	//}

	atleast_one_odd := false
	N := e - s + 1
	//mincount := N + 1
	big_is_min := N
	min_left := N
	min_right := 0
	//var odd_left int
	//var odd_right int
	for i := 0; i < 36; i++ {
		v := arr[e+1][i] - arr[s][i]

		if v%2 != 0 {
			atleast_one_odd = true
			odd_left, odd_right := find_odd_left_right(i, s, e, arr, arr_pos)
			//fmt.Printf("odd_left: %v, odd_right: %v\n", odd_left, odd_right)
			right_part_size := e - odd_left
			left_part_size := odd_right - s
			bigger_part := int(math.Max(float64(left_part_size), float64(right_part_size)))
			if bigger_part < big_is_min {
				big_is_min = bigger_part
				//mincount = v
				//minch = k
				min_left = odd_left
				min_right = odd_right
			}
		}
	}

	if !atleast_one_odd {
		//ab := &AB{s, e}
		//dp_map[*ab] = e - s + 1
		//fmt.Printf("putting in db_map: %v\n", dp_map)
		ret := e - s + 1
		if ret > max_len {
			max_len = ret
		}
		return ret
	}

	ret := 0

	//find the left most character position of minch
	// and recusively invoke the function for the right part ignoring that character
	//fmt.Printf("recursive invoke for right part of i: %v\n", i)
	rbeg := min_left + 1
	rend := e
	l1 := rend - rbeg + 1
	//fmt.Printf("rbeg: %v, rend: %v\n", rbeg, rend)

	// now take the right most minch position
	// and invoke it for the left part
	//fmt.Printf("recursive invoke for left part of i: %v\n", i)
	lbeg := s
	lend := min_right - 1
	l2 := lend - lbeg + 1
	//fmt.Printf("lbeg: %v, lend: %v\n", lbeg, lend)
	//fmt.Printf("---\n")

	// We invoke on the longer one first and then invoke on the other
	// only if returned lenght is shorter than the other
	if l1 < l2 && l2 > max_len {
		tmp := recurse_fn(S, lbeg, lend, arr, arr_pos)
		if tmp > ret {
			ret = tmp
		}

		if ret < l1 && l1 > max_len {
			tmp = recurse_fn(S, rbeg, rend, arr, arr_pos)

			if tmp > ret {
				ret = tmp
			}
		}
	} else if l1 > max_len {
		tmp := recurse_fn(S, rbeg, rend, arr, arr_pos)
		if tmp > ret {
			ret = tmp
		}

		if ret < l2 && l2 > max_len {
			tmp = recurse_fn(S, lbeg, lend, arr, arr_pos)

			if tmp > ret {
				ret = tmp
			}
		}
	}

	//ab = &AB{s, e}
	//dp_map[*ab] = ret
	//fmt.Printf("bottom putting in db_map: %v\n", dp_map)
	if ret > max_len {
		max_len = ret
	}
	return ret
}

func find_odd_left_right(ch, s, e int, arr, arr_pos [][]int) (int, int) {

	pos_list := arr_pos[ch]
	//fmt.Printf("pos_list: %v\n", pos_list)
	//tot_n := len(pos_list)
	before_n := arr[s][ch]
	last_n := arr[e+1][ch]

	left := pos_list[before_n]
	right := pos_list[last_n-1]
	return left, right
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

func main() {
	S := "yyyyy"
	ret := Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)

	S = "bxdcbadcbaya"
	ret = Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)

	S = "22"
	ret = Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)

	S = "6daa6ccaaa6d"
	ret = Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)

	S = "abca"
	ret = Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)

	S = "gggaaaacdbbbbxeffffff"
	ret = Solution(S)
	fmt.Printf("S: %v; ret: %v\n", S, ret)
}
