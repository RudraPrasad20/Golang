package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {

	// Arrays

	var arr [3]int        // create an array of int type of length 3
	fmt.Println(len(arr)) // length

	arr[2] = 1          // add 1 at index 2
	fmt.Println(arr[1]) // print index 1

	// print whole arr
	// if arr is empty, it will print 0 (default)
	fmt.Println(arr) // [0 0 1]

	var vals [4]bool  // create an array of bool type of length 3
	vals[2] = true    // 2nd index - true
	fmt.Println(vals) // false by default

	// similarly it will print empty strings in case of string array
	// int -> 0, string →› ', bool -> false)

	num := [3]int{1, 2} // [1 2 0]
	fmt.Println(num)

	// 2d array
	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)

	// Slices =>

	// uninitialized slice is nil - null in ts
	var nu []int
	fmt.Println(nu)        // []
	fmt.Println(nu == nil) // true
	fmt.Println(len(nu))   // 0

	// length 0 and capacity 6
	var tp = make([]int, 2, 6) // 2 - initial len | 6 - max capacity
	// add numbers at the end of an array
	tp = append(tp, 3) // [0 0 3]
	tp = append(tp, 6) // [0 0 3 6]
	tp = append(tp, 8) // [0 0 3 6 8]
	fmt.Println(tp)    // [0 0]
	// if we add more larger then the capacity. it will double that capacity, because array is dynamic
	// here the max capacity is 6 but after adding few more integers it will change to 12 - multiply of the capacity - 6x2
	fmt.Println(cap(tp))   // capacity - 6
	fmt.Println(tp == nil) // false

	var n = make([]int, 3, 5)
	n[2] = 4
	n[1] = 5
	n = append(n, 1)
	fmt.Println(n)
	fmt.Println(cap(n))
	fmt.Println(len(n))

	// copy fn
	var num1 = make([]int, 0, 5)
	num1 = append(num1, 3)
	var nums2 = make([]int, len(num1))

	// copy
	// where to paste, what to copy - (destination, src)
	copy(nums2, num1)
	fmt.Println(num1, nums2)

	// slice operator

	var ar = []int{1, 2, 3}
	// index
	fmt.Println(ar[0:2]) // [1,2]
	fmt.Println(ar[:3])  // [1,2,3]
	fmt.Println(ar[1:])  // [2,3]

	var numi1 = []int{1, 2, 3}
	var numi2 = []int{1, 2, 4}
	// returns true or false
	fmt.Println(slices.Equal(numi1, numi2)) // false - because those 2 arrays are not same

	var rums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(rums)

	// MAPS:

	// create a map
	m := make(map[string]string)
	// setting an element
	m["name"] = "go"
	m["type"] = "strr"

	// m := make(map[string]int)
	// m["class"] = 10

	// enter the key to get the element
	// IMP: if key does not exists in the map then it returns zero
	fmt.Println(m["name"], m["type"])
	fmt.Println(len(m)) // 2

	// delete
	delete(m, "type") // delete the type keyword line
	fmt.Println(m)    // go

	// clear
	clear(m) // clear the map - remove all in map

	// m := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(m) // map [phones: 3 price: 4Q]

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"paid": 4000, "got": 5000}
	// it only works with this type:  m := map[string]int{"price": 40, "phones": 3}
	// to use this, we have to declear the map and declear things in a single line
	va, ohh := m2["paid"]
	fmt.Println(va)
	if ohh {
		fmt.Println("all right")
	} else {
		fmt.Println("wrong")
	}

	fmt.Println(maps.Equal(m1, m2)) // true or false

	// range

	numb := []int{6, 7, 8}
	total := 0

	for index, TotalNum := range numb {
		fmt.Println(TotalNum, index)
		total = total + TotalNum
	}
	fmt.Println(total)

	m3 := map[string]string{"fname": "dana", "Iname": "lan"}
	// print key and value
	for key, val := range m3 {
		fmt.Println(key, val)
	}
	// can print only key but cant print only value
	for ky := range m3 {
		fmt.Println(ky)
	}
	// unicode code point runeb -> similar to ascii
	// starting byte of rune
	for i, c := range "rana" {
		fmt.Println(i, c, string(c))
	}
}
