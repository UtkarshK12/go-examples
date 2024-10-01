package main

import (
	"fmt"
	"strconv"
)

//& (Address-of operator):
// Used to get the memory address of a variable.
// Returns a pointer to the variable.


// * (Dereference operator):

// Used in two contexts:
// a. To declare a pointer type.
// b. To access the value stored at a pointer's address.


func main(){
	a:=5
	b:=10
	swapValues(a,b)
	//nothing happens here
	fmt.Println("a->"+strconv.Itoa(a)+" b->"+strconv.Itoa(b))

	swapPointers(&a,&b)
	//the values are now reversed
	fmt.Println("a->"+strconv.Itoa(a)+" b->"+strconv.Itoa(b))

	//Slices and Maps: Although passed by value, the copy includes a pointer to the underlying data.

	//Modifications affect the original data, but reassigning the slice/map variable inside the function doesn't affect the original variable.
	//However, operations that change the slice itself (like append that causes reallocation) won't affect the original slice variable.

	fmt.Println("working with slices")

	var s []int
	s = append(s, 3)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 6)

	appendSlice(s)
	//the original slice is not modified because append creates a new slice and creates a new pointer from address "s" to new slice
	//but the old slice remains untouced
	fmt.Println(s)

	modifySlice(s)
		//if we try to modify the values in slice, it does not create a new slice but modifies the value in place ... so in this case the base slice gets modified
	fmt.Println(s)

}

func appendSlice(s []int) {
	s = append(s, 8)
}

func modifySlice(s []int){
	s[0]=100
}

func swapValues(a int, b int){
	c:=b
	b=a
	a=c
}

func swapPointers(a *int, b *int){
	c:=*b
	*b=*a
	*a=c
}