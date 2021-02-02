package main

import (
    "fmt"
    "math"
    "runtime"
)

func sqrt(x int) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }

    return fmt.Sprint(math.Sqrt(float64(x)))
}

func powerlimit(x, y, lim float64) float64 {
    /* The scope of v is limited to the if/else block */
    if v := math.Pow(x, y); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func NewtonianSqrt(x float64) float64 {
    z := 1.0
    count := 1
    for {
        fmt.Printf("Guess number %d is %f\n", count, z)
        if math.Abs(z*z - x) < 0.000001 {
            return z
        }
        z -= (z*z - x) / (2*z)
        count += 1
    }
}

func runningOS() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        fmt.Printf("%s.\n", os)
    }
}

func main() {

    defer fmt.Println("Finishing. Goodbye!")
    defer runningOS()

    sum := 0

    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    /* While loop */
    sum = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)

    /* 
    * infinite loop
    for {

    }
    */

    /* Both calls to pow return their results before the call to fmt.Println in main begins. */
    fmt.Println(sqrt(2), sqrt(-4))
    fmt.Println(
        powerlimit(3, 2, 10),
        powerlimit(3, 3, 20),
    )

    x := 20.0
    fmt.Println(NewtonianSqrt(x))
    fmt.Println(math.Sqrt(x))
}
