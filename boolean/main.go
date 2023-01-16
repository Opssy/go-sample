package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age >= 50)

	if age > 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is greater than 40")
	} else {
		fmt.Println("age is not less than 45")
	}
	//slicing method
	name := []string{"mario", "yoshi", "luigi", "opssy"}

	for index, value := range name {
		if index == 1 {
			fmt.Println("continuing at pos", index)
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
