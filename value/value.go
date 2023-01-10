package main

	
import (
    "fmt"
    "math"
)

const s string = "constant"
func main(){
	// fmt.Println("go" + "lang")
	// fmt.Println(1+1)
	// fmt.Println(7.0/3.0 )
	// fmt.Println(true && false)
    // fmt.Println(true || false)
    // fmt.Println(!true)

	// var a  = "initial"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b,c)

	// f := "apple"
	// fmt.Println(f)

	//constant
    fmt.Println(s)

	const n = 500000000
    
	const d = 3e20 /n 
	fmt.Println(d)
    
	fmt.Println(math.Sin(n))

	// for loop

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	
	for {
		fmt.Println("loop")
		break
	} 
	for n :=0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
