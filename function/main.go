package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}
func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func main() {
	sayGreetings("Ope")
	sayBye("dayo")
	cycleName([]string{"luigi", "cloud", "ope", "dmj"}, sayGreetings)
}
