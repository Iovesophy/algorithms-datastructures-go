package zeroone

import "fmt"

func Sum() {
	n := 10
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i
	}
	fmt.Println(s)
}
