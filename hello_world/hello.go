package main

import (
	"fmt"
	"math"
	"math/rand"
)

func swap(x, y string) (string, string) {
	return y, x
}

var i, j int = 1, 2

func main() {
	fmt.Println(swap("Hello", "世界"))
   fmt.Println("My favorite number is", rand.Intn(100))
   fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
   fmt.Printf("Pi is %g\n", math.Pi)

   var c, python, java = true, false, "no!"
   fmt.Println(i, j, c, python, java)
}
