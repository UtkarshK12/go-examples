package main
import "fmt"

//method is basically a property of the struct while a function takes a struct as an argument
type rectangle struct{
	a int 
	b int
}

//this is a function
func calculateArea(rect *rectangle) int{
	return rect.a*rect.b;
}

//this is a method
func (rect rectangle) calaculateArea() int{
	return rect.a*rect.b
}

func StructsVsMethods(){
	fmt.Println("inside structs vs methods")

	//calling a function
	rect:=rectangle{a:10, b:20}
	fmt.Println(calculateArea(&rect))

	//calling a method
	rect2:=rectangle{a:20, b:40}
	fmt.Println(rect2.calaculateArea())
}