package main

import (
	"fmt"
	"slices"
)
		

func SliceDemo(){

	fmt.Println("Inside slices")

	//declaring a slice with fixed length
	a:=make([]int,3)
	a= append(a, 1)
	fmt.Println(a)

	//declaring a slice with variable length
	var b []int;
	b = append(b, 1)
	b = append(b, 3)

	fmt.Println(b)


	//array vs slice declaration
	// array - a:=[5]int{} slice - a:=[]int{}
	//array -> var a = [...]int{1,2,3}   slice -> var a= []int{};
	
	c:=[]int{0,0,0,1}
	if slices.Equal(a,c){
		fmt.Println("same")
	}
	


}