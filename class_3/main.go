package main

import "fmt"

func add(a int, b int, c, d string) int {
	fmt.Println(a + b)
	fmt.Println(c, d)
	return a + b
}

func pn() (string, string, string, int) {
	return "heehe", "hihi", "hell", 30
}

// variadic fn:
// by doing this, we can only pass integers, we cant pass string or any other types
// check where we called in the main function for more clarity
// (num ...interface{}) -> for any type
func vard(num ...int) {
		fmt.Println(num)
	}

func main() {
	add(3, 4, "hell", "yeh")
	fmt.Println(pn())

	// variadic fn:
	 vard(39,3,3,3,3)

	 // CLOSURES

}
