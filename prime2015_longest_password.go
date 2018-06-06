package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "strings"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4

	words := strings.Split(S, " ")

	ret := -1
	for _, word := range words {
		N := len(word)
		cc := 0
		nc := 0
		bad := false
		for i := 0; i < N; i++ {
			ch := word[i]
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				cc++
			} else if ch >= '0' && ch <= '9' {
				nc++
			} else {
				bad = true
				break
			}
		}

		if !bad && cc%2 == 0 && nc%2 != 0 {
			if ret == -1 || ret < N {
				ret = N
			}
		}

	}

	return ret
}
