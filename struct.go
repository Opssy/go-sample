package main
import (
	"math"
    "fmt"
)

type shape interface{
 area() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
func(r rect) area() float64{
	return r.width * r.height
}
func (c circle) area() float64{
	return math.Pi * c.radius*c.radius
}
func printArea(sh shape) {
	fmt.Printf("%v area is %v\n", sh, sh.area())
}
func main(){
	r := rect{
     width: 3,
	 height: 5,
	}
	c := circle{
		radius: 5,
	}
	printArea(r)
	printArea(c)
}