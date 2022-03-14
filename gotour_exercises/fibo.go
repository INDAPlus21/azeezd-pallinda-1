package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// Store current and previous values, changing them on call
	i0, i1 := 0, 1
	return func() int {
		val := i0
		i0, i1 = i1, val+i1

		return val
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
