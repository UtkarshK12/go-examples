package main

import (
	"fmt"
	"math/rand"

)
func Closures(){
	fmt.Println("inside closures now")
	fetchFunc:=apiFetch()

	fmt.Println(fetchFunc())
	fmt.Println(fetchFunc())
	fmt.Println(fetchFunc())
}

func apiFetch() func()int {
	//closures return a function instead of a value
	return func () int{
		return rand.Intn(10)
	}
}

//Function Factories:
// Closures can be used to create functions with pre-set parameters.

// Closure as Function Factory

// Callbacks and Event Handlers:
// Closures are excellent for creating callback functions that need to access data from their outer scope.

// Closure as Callback

// Memoization and Caching:
// Closures can help implement memoization for expensive function calls.

// Closure for Memoization

// Middleware in Web Applications:
// Closures are commonly used to create middleware functions in web frameworks.

// Closure as Middleware

// Implementing Iterators:
// Closures can be used to create stateful iterators.

// Closure as Iterator

// Partial Function Application:
// Closures allow you to partially apply a function, creating a new function with some parameters pre-set.

// Closure for Partial Application

// Encapsulation and Data Hiding:
// Closures can be used to create private variables and methods.

// Closure for Encapsulation