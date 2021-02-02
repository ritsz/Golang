package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        ret := a
        a = b
        b = ret + a
        return ret
    }
}

func main() {
    pos, neg, fib := adder(), adder(), fibonacci()
    /* Each closure is bound to its own sum variable.  */
    for i := 0; i < 20; i++ {
        fmt.Println(
            pos(i),
            neg(-i),
        )
        fmt.Println(fib())
    }
}

