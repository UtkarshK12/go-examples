package main

import ("fmt")

func VariadicFn(){
	fmt.Println("inside variadic function")
	res1:=variadic_sum(1,2,3,4,5)
	fmt.Println(res1)

	res2:=variadic_sum(1,2,3,4,5,6,7,8)
	fmt.Println(res2)
}


func variadic_sum(nums ...int) int {
	//the arguments to variadic functions become slices in go
	fmt.Println(nums)
	sum:=0
	for _,num:= range nums {
		sum+=num
	}
	return sum
}



