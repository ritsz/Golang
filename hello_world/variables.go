package main

import "fmt"

var i, j int
var a, b = 1, "Hello"
/*
 Non declaratibe statement outside function body is an error
 a, b := 1, "Hello" is an Error"

 Outside a function every statement begins with a keyword like var, func etc.
*/


var (
   bl       bool     = false
   MaxInt   uint64   = 1<<64 - 1
   str      string   = "Goodbye"
)

func main() {
    var c, python = true, false
    i, j := 10, 20 /* Inside a func, variables can be non declarative. */
    a, b = 9, "World" /* Updating values with =, := only for new variables */
    fmt.Println(a, b, c, python, i, j)

    fmt.Printf("Type: %T Value: %v\n", bl, bl)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", str, str)
}
