package main

import "fmt"

func main(){
	a:=5
	b:=6
	c:=7
	res:=sum(a,b)
	fmt.Println(res)

	res2:=sum2(a,b,c)
	fmt.Println(res2)

	res3, res4:=sumAndMult(2,3)
	fmt.Println(res3, res4)
}


func sum(a int, b int) int {
	return a+b
}

// go does not support function overloading so the function name needs to be different
func sum2(a int , b int , c int) int {
	return a+b+c;
}

//multiple return values, nice feature in go
func sumAndMult(a int, b int ) (int , int) {
	return a+b,a*b
}