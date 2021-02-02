package main

import "fmt"

/*
  * An empty interface may hold values of any type. (Every type implements at least zero methods.)
  * Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}. 
*/

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
   /*
    * Calling interface function on nill interface will cause SIGSEGV
    * Calling a method on a nil interface is a run-time error because there is 
    * no type inside the interface tuple to indicate which concrete method to call. 
    *
    * i.M() <- SIGSEGV here
    */

   /* Using interface with nil underlying value. */
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

