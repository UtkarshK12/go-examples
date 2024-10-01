package main

import "fmt"

func Constants(){
	fmt.Println("Inside constants")
	//changing this from 4.0 to 4 will change the output to 4 as the type casting does not happen
	const a =4.0 
	fmt.Println(a)

	const d = a/10
	fmt.Println(float64(d))
}