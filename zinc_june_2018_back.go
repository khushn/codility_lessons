/**
Kate was given a birthday gift of three theater tickets. Now she is browsing the theater program for the next N days. On the program, performances are named by integers. Every day, one performance is staged. Kate wants to choose three days (not necessarily consecutive) to go to the theater.

In how many ways can she use her tickets? Two ways are different if the sequences of watched performances are different. Kate likes the theater, so she may watch one performance more than once. For example, if N = 4 and theater program looks as following: [1, 2, 1, 1], Kate has four possibilities to choose the dates: [1, 2, 1, 1], [1, 2, 1, 1], [1, 2, 1, 1], and [1, 2, 1, 1], but they create only three different sequences: (1, 2, 1), (1, 1, 1) and (2, 1, 1). The correct answer for this example is 3. Notice that the order of performances matters, so the first and the last sequences are considered different.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, denoting names of performances for the next N days, returns the number of possible ways to spend the tickets. Since the answer can be very large, provide it modulo 109 + 7 (1,000,000,007).

For example, given A = [1, 2, 1, 1], the function should return 3 as exmplained above.

Given A = [1, 2, 3, 4], the function should return 4. There are four ways to spend tickets: (1, 2, 3), (1, 2, 4), (1, 3, 4) and (2, 3, 4).

Given A = [2, 2, 2, 2], the function should return 1. There is only one way to spend tickets: (2, 2, 2).

Given A = [2, 2, 1, 2, 2], the function should return 4. There are four ways to spend tickets: (1, 2, 2), (2, 1, 2), (2, 2, 1) and (2, 2, 2).

Given A = [1, 2], the function should return 0. Kate cannot use all three tickets in only two days.

Assume that:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [1..N].
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
Custom test casesformat: [1, 2, 1, 1]	0/10

**/

/**
STATUS:

Incomplete. This is a tough one.

Update on 7/June/2018. Got an idea. Works partially:

Results:
https://app.codility.com/cert/view/certBNKWBT-7PVXNJANNEVDN8T5/details/

**/

package main

//package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

const MAX = 1000000007

func Solution(A []int) int {
	// write your code in Go 1.4

	N := len(A)
	if N < 3 {
		return 0
	}

	// num uniques
	nq := 0

	// uniques count

	// left to right
	arr1 := make([]int, N)

	// counting occurences
	map1 := make(map[int]int)
	// map to pot positions of a program
	map2 := make(map[int][]int)
	for i := 0; i < N; i++ {
		a := A[i]
		x, ok := map1[a]
		if !ok {
			map1[a] = 1
			arr := make([]int, 1)
			arr[0] = i
			map2[a] = arr

			// uniques counter
			nq++
			arr1[i] = nq
		} else {
			map1[a] = x + 1
			arr := map2[a]
			arr = append(arr, i)
			map2[a] = arr
			arr1[i] = arr1[i-1]
		}
	}

	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("arr1: %v\n", arr1)

	// right to left
	arr2 := make([]int, N)
	tc := 0
	tmap := make(map[int]int)
	for i := N - 1; i >= 0; i-- {
		a := A[i]
		x, ok := tmap[a]
		if !ok {
			tmap[a] = 1
			tc++
			arr2[i] = tc
		} else {
			tmap[a] = x + 1
			arr2[i] = arr2[i+1]
		}
	}

	fmt.Printf("arr2: %v\n", arr2)

	ret := 0
	// counting 3s like '111'
	for _, v := range map1 {
		if v >= 3 {
			ret++
		}
	}

	//fmt.Printf("after 3s count: %v\n", ret)
	//fmt.Printf("nq: %v\n", nq)
	//fmt.Printf("map2 size: %v\n", len(map2))

	// counting 2 repitions like 33x, 3x3 and x33
	for k, arr := range map2 {

		fmt.Printf("k: %v\n", k)
		if len(arr) >= 2 {

			// first -- 33x
			ind := arr[1] + 1
			if ind < N {
				p := arr2[ind]
				// uniques after this pos
				fmt.Printf("33x -> p: %v\n", p)
				u := nq - p
				ret += u % MAX
			}

			// second -- 3x3
			p1 := arr1[arr[0]]
			pe := arr1[arr[len(arr)-1]]
			fmt.Printf("3x3 --> p1: %v, pe: %v\n", p1, pe)
			u := pe - p1
			ret += u % MAX

			// third -- x33
			ind = arr[len(arr)-2] - 1
			if ind >= 0 {
				p2 := arr1[ind]
				fmt.Printf("x33 -> p2: %v\n", p2)
				u = p2 - 1
				ret += u % MAX
			}
		}

	}

	//fmt.Printf("after 2s count: %v\n", ret)

	// counting no reps e.g. 123, 456 and so on
	ret += ncx(nq, 3) % MAX

	return ret
}

func ncx(n, x int) int {

	if n < x {
		return 0
	}

	if x == n {
		return 1
	}

	xfact := 1
	for i := 2; i <= x; i++ {
		xfact *= i
	}

	//fmt.Printf("xfact: %v\n", xfact)

	ret := 1
	for i := n - x + 1; i <= n; i++ {
		ret *= i
		ret = ret % MAX
	}

	//fmt.Printf("inside factorial ret: %v\n", ret)

	ret /= xfact

	return ret
}

func main() {

	var atest = []int{1, 2, 1, 2}
	ret := Solution(atest)
	fmt.Printf("ret: %v\n", ret)
}
