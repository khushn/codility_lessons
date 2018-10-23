/**

**/

/*Results

Easy problem. Golden award in 1st try. Around 20 minutes.
https://app.codility.com/cert/view/certEJE7YY-KBAM47FHBMHSV6PV/

*/

package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, M int, X []int, Y []int) int {
	// write your code in Go 1.4
	K := len(X)

	if K%2 != 0 {
		return 0
	}

	x := make([]int, N)
	y := make([]int, M)

	for i := 0; i < K; i++ {

		x[X[i]] += 1
		y[Y[i]] += 1

	}

	fmt.Printf("x := %v\n", x)
	fmt.Printf("y := %v\n", y)

	retx := 0

	xc := 0

	for i := 0; i < N; i++ {
		xc += x[i]
		if xc == K/2 {
			retx += 1
		}

	}

	rety := 0
	yc := 0
	for i := 0; i < M; i++ {
		yc += y[i]
		if yc == K/2 {
			rety += 1
		}
	}

	return retx + rety
}
