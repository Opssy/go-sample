package main 

import (
	"fmt"
)

func main(){
	menu := map[string]float64{
		"soup":  4.99,
		"pie": 5.0,
		"salad": 6.99,
		"toffe": 3.55,

	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])


	//loop
	for k, v := range menu{
		fmt.Println(k, "_", v)
	}
}