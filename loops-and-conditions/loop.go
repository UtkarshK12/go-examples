package main 

import "fmt"

func main(){
	
	//for loop
	for i:=0;i<10;i++{
		fmt.Println(i)
	}

	//using range
	fmt.Println("using range")

	for i:= range 4{
		fmt.Println(i)
	}

	//looping through array
	fmt.Println("array loops")
	c:=[]int{1,2,3,4,5}
	for i:=range c {
		fmt.Println(i)
	}


	//apart from function calling, no such () brackets found yet
	fmt.Println("using if else")
	a:=5
	if a<6 {
		fmt.Println("lesser than 6")
	}else{
		//the else statement should be in the same line where if statement ends ??
		fmt.Println("greater than 6")
	}
}