/**
Joe the Farmer owns a square plot of farmland whose sides are of length N. The land is split into N rows and N columns of equal size, so that there are N2 identical square fields within the farmland. Thanks to the abundant crops that Joe harvested last year, he could afford to buy N sprinklers. They significantly reduce the amount of time that Joe spends watering his plants.

Every sprinkler can be placed in any field of the farmland, but no two sprinklers can occupy the same field. Upon activation, each sprinkler irrigates every field within the same column and row in which it appears.

After delivery, the sprinklers were placed in a batch, so the K-th sprinkler is positioned in field (XK, YK). Joe knows that this arrangement is not optimal, as some fields may not be watered by any sprinkler. Moreover, no column or row should be watered by more than one sprinkler, as it may cause over-hydration of the crops.

In one move, the farmer can move a single sprinkler to an adjacent unoccupied field. Two fields are adjacent to one another if they have a common side.

What is the minimum number of moves required to rearrange the sprinklers so that all fields will be irrigated by sprinklers, and no two sprinklers will water the same column or row? Since the answer can be very large, provide it modulo 1,000,000,007 (109 + 7).

Write a function:

func Solution(X []int, Y []int) int

that, given arrays X and Y, both consisting of N integers and representing the coordinates of consecutive sprinklers, returns a minimal number of moves modulo 1,000,000,007, required to irrigate all fields.

For example, given array X = [1, 2, 2, 3, 4] and array Y = [1, 1, 4, 5, 4] the function should return 5, as we can make following moves:



For another example, given array X = [1, 1, 1, 1] and array Y = [1, 2, 3, 4] the function should return 6:



Given array X = [1, 1, 2] and array Y = [1, 2, 1] the function should return 4:



Assume that:

N is an integer within the range [2..100,000];
each element of arrays X, Y is an integer within the range [1..N];
each sprinkler appears in a distinct field (no field may contain more than one sprinkler).
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N) (not counting the storage required for input arguments).
Copyright 2009–2018 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
**/

/**Result

Got the golden award.
https://app.codility.com/cert/view/certNERVBZ-WUCYXZMSSY7JJDYC/

Before that this is one of the unsuccessful attempts
https://app.codility.com/cert/view/certH4YTB2-JV2SVRKU5MCVZ34R/details/
(Note how complicated the code got. Simplicty rules!!)

*/

package main

// package solution

// you can also use imports, for example:
import "fmt"

// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X []int, Y []int) int {
	// write your code in Go 1.4

	const MAX = 1000000007

	N := len(X)
	x_arr := make([]int, N)
	y_arr := make([]int, N)
	for i := 0; i < N; i++ {
		// making them 0-indexed
		x := X[i] - 1
		y := Y[i] - 1

		// populate the actual positions
		x_arr[x] += 1
		y_arr[y] += 1
	}
	//fmt.Printf("x_arr: %v\ny_arr: %v\n", x_arr, y_arr)

	ret := 0
	xds := 0
	yds := 0

	for i := 0; i < N; i++ {

		// first we consider along X-axis
		xd := x_arr[i] - 1
		xds += xd
		ret += int(math.Abs(float64(xds)))
		ret %= MAX

		// I realize, can do the same thing for y-axis as well
		// since they are independent
		yd := y_arr[i] - 1
		yds += yd
		ret += int(math.Abs(float64(yds)))
		ret %= MAX

	}

	return ret
}

func main() {
	X := []int{1, 2, 2, 3, 4}
	Y := []int{1, 1, 4, 5, 4}
	fmt.Printf("---------\nX: %v\nY: %v\n", X, Y)
	ret := Solution(X, Y)
	fmt.Printf("ret: %v\n", ret)

	X = []int{1, 1, 1, 1}
	Y = []int{1, 2, 3, 4}
	fmt.Printf("---------\nX: %v\nY: %v\n", X, Y)
	ret = Solution(X, Y)
	fmt.Printf("ret: %v\n", ret)

	X = []int{1, 1, 2}
	Y = []int{1, 2, 1}
	fmt.Printf("---------\nX: %v\nY: %v\n", X, Y)
	ret = Solution(X, Y)
	fmt.Printf("ret: %v\n", ret)
}
