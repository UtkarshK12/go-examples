package main
import "fmt"

//shorthand declaration outside function is prohibited
//c:=4


//public functions in go which can be accessed from other functions start with caps
func Variables(){
	fmt.Println("Inside variables")

	var a=4;
	//short hand method
	b:=4

	fmt.Println(a+b)
}
//go run * to run all files in a package