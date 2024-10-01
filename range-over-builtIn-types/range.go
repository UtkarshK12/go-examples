package main
import ("fmt"
		"strconv")

func main(){
	a:=[]int{1,2,3,4,5}

	//we can leave the first value as blank as well
	for index,value:=range a{
		fmt.Println(strconv.Itoa(index)+" "+strconv.Itoa(value))
	}

	//maps
	m:=make(map[string]int)
	m["key1"]=5
	m["key2"]=7
	m["key3"]=4
	m["key4"]=8

	for index,value:=range m{
		fmt.Println(index+" "+strconv.Itoa(value))
	}

}