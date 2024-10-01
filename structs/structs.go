package main 

import "fmt"


//declaring a struct
type person struct{
	name string
	age int
}

func sayHi(p *person){
	fmt.Println("Hi "+ p.name)
}

func newPerson(name string, age int) person{
	return person{name: name, age: age}
}


func main(){
	fmt.Println("Inside struct")

	//initializing a struct
	person1 := person{name : "uto", age:24};
	fmt.Println(person1)


	//calling a function
	sayHi(&person1)

	//creating a person using function 
	person2:=newPerson("uto",25)
	fmt.Println(person2)

	StructsVsMethods()
	

}