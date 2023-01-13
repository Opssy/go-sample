package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
greeting := "hello there friends"

fmt.Println(strings.Contains(greeting, "hello"))
fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
fmt.Println(strings.ToUpper(greeting))
fmt.Println(strings.Index(greeting,"ll"))

age := []int{45, 20, 30, 75, 60, 50, 23}

sort.Ints(age)
fmt.Println(age)

index := sort.SearchInts(age, 30)
fmt.Println(index)

name := []string{"yoshi", "mario", "peach", "bower", "luigi"}

sort.Strings(name)
fmt.Println(name)
}