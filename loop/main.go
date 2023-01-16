package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}
	//how to loop a slice of string

	name := []string{"mario", "yoshi", "luigi", "opssy"}
	//for i := 0; i <= len(name); i++ {
	//	fmt.Println(name[i])
	for index, value := range name {
		fmt.Printf("the value at index %v is %v", index, value)
	}
}
