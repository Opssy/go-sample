package calculator_app

import (
	"flag"
	"fmt"
	"os"
)

func Addition(a int, b int) int {
	return a + b
}
func main() {
	add := flag.Bool("add", false, "Add two numbers together")
	flag.Parse()

	var a, b int
	fmt.Println("Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Enter 2nd Number: ")
	fmt.Scan(&b)

	switch {
	case *add:
		fmt.Printf("Additon: %d \n", Addition(a, b))
	case *subs:
		fmt.Printf("Difference: %d \n", subtract(a, b))
	case *mult:
		fmt.Printf("Product: %d \n", multiply(a, b))
	case *div:
		fmt.Printf("Division: %d \n", division(a, b))
	default:
		fmt.Fprintln(os.Stderr, "Wrong option try with add, subtract, div and multply")
		os.Exit(1)
	}
}
