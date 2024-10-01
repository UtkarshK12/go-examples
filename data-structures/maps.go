package main

import (
	"fmt"
	
)

func Maps(){

	fmt.Println("inside maps")
	//creating a map
	m:=make(map[string]int)

	//adding values to a map
	m["key1"]=3
	m["key2"]=5

	//retrieving values
	count:=m["key1"]
	fmt.Println(count)

	//optional second return value is a bool to see if the value exists in a map
	count,check:=m["key3"]
	fmt.Println(check)


	//creating a map with some values (make not needed here)
	m2:=map[string]int{"foo":1,"bar":2}
	fmt.Println(m2)
}