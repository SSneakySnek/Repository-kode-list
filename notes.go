// Always save when you want to run in terminal
package main

import "fmt"

func notes() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "bruh"
	nameThree = "link"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "herro" // cant use outside of a function
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	//float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 342342342.3
	scoreThree := 1.5

	age := 20
	name := "Thomas"

	//Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n ")

	//Println
	fmt.Println("Hello bruh")
	fmt.Println("bye bruh")
	fmt.Println("I am", age, "old, and my name is", name)

	//Printf (formatted strings) &_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points \n", 226.32)
	fmt.Printf("you scored %0.1f points \n", 226.32)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi" // position 1 = mario cause yoshi is 0

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices  (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 86)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "bruh")
	fmt.Println(rangeOne)
}
