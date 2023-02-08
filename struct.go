package main

import "fmt"

//
//type shape interface {
//	area() float64
//}
//
//type rect struct {
//	width, height float64
//}
//type circle struct {
//	radius float64
//}
//
//func (r rect) area() float64 {
//	return r.width * r.height
//}
//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//func printArea(sh shape) {
//	fmt.Printf("%v area is %v\n", sh, sh.area())
//}
//func main() {
//	r := rect{
//		width:  3,
//		height: 5,
//	}
//	c := circle{
//		radius: 5,
//	}
//	printArea(r)
//	printArea(c)
//}

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {

	engineer := Engineer{
		Name: "Elliot",
		Age:  27,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}
	fmt.Printf("%+v\n", engineer)

	fmt.Println(engineer.Project.Name)

}
