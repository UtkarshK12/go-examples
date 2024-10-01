package main

import "fmt"

func main (){

	//variable sizearray declaration
	a:=[...]int{1,2,3,4,5}
	fmt.Println(a)

	var c=[...]int{1,2,3,4}
	fmt.Println(c)

	//declaring array array
	//after creating an array, the size does not change. So to fix the size at the start
	b:=[4]int{}
	fmt.Println(b)

	b=[...]int{1,2,3,4}
	fmt.Println(b)

	//2D array
	twoD:=[3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(twoD)

	SliceDemo()
	Maps()
}