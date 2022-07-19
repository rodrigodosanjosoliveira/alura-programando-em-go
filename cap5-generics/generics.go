package main

import "fmt"

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
}

// T is a type parameter that is used like normal type inside the function
// any is a constrant on type i.e T has to implement "any" interface
func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}

	return r
}
