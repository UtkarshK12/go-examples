package main

import "fmt"


//in go interfaces are a collection of method signatures. ONLY METHODS
type shape interface{
	area() int
} 

type rect struct{
	a int 
	b int
}

type circle struct{
	r int
}

func (r rect) area() int {
	return r.a*r.b
}

func (c circle) area() int {
	return 3*c.r*c.r
}

func printArea(s shape){
	fmt.Println(s.area())
}

//Note that unlike some (most) languages, your structs don't explicity declare that they implement the interface. You don't write code like this.

func main(){
	fmt.Println("inside interfaces")

	r:=rect{5,10}
	c:=circle{5}

	printArea(r)
	printArea(c)
}