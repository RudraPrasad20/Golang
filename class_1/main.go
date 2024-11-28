package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hey less go..")

	// var price float32 = 50.5
	// var price = 50.5
	price := 50.5 // but we cant declear this syntax outside the function
	fmt.Println(price)

	// declear multiple constants
	const (
		age  = 30
		name = "raman"
	)
	fmt.Println(age, name)

	// loop:
	// while loop:
	var in = 1
	for in <= 3 {
		fmt.Println(in)
		in = in + 1
	}

	// infinite loop:
	// for {
	// 	fmt.Println("ok")
	// }

	// for loop:
	for i := 0; i <= 3; i++ {
		// break
		// if i == 2 {
		// 	continue
		// 	// it will skip number 2
		// }
		fmt.Println(i)
	}

	// range:
	// another way to write the same for loop:
	for im := range 5 {
		// print from 1 to 4
		fmt.Println(im)
	}

	// if
	var class = 6
	if class >= 4 {
		fmt.Print("big dawww")
	} else if class >= 8 {
		fmt.Print("mid daww")
	} else {
		fmt.Print("small dawww")
	}

	var isKid = true
	var marks = 80
	// or - one of the condition has to be true
	if isKid || marks == 60 {
		fmt.Println("good to go")
	}
	// and - both the conditions has to be true
	if isKid && marks == 60 {
		fmt.Println("good to go")
	}

	// we can declare a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age > 12 {
		fmt.Println("person is teenager", age)
	}

	// simple switch
	int := 20
	switch int {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	default:
		fmt.Print("other")
	}

	// multiple switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("chutti he enjoy..")
	default:
		fmt.Print("aa jao sare kaam pe")
	}

	// type switch
	// interface{} -> any type in ts
	iam := func(i interface{}) {
		switch ty := i.(type) {
		case bool:
			fmt.Print("its an bool")
		case string:
			fmt.Print("its a str")
		default:
			fmt.Print("other type", ty)
		}
	}
	iam(true)

}
