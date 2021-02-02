package main

import (
    "fmt"
    "math"
)

/*
* Float type and its functions
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

/*
* Vertex type and its functions
*/
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

type Line struct {
   vertex1 Vertex
   vertex2 Vertex
}

func (l Line) Abs() float64 {
    return math.Abs(l.vertex1.Abs() - l.vertex2.Abs())
}

/* Abser acts as an interface type for any type that implements Abs() float64 function
* Abser can acts as an interface to Vertex and MyFloat since both of them implement atleast the Abs() function.
*/
type Abser interface {
    Abs() float64
}

/* Scaler interface for a type that implements 2 functions type ATLEAST. */
type Scaler interface {
    Abs() float64
    Scale(float64)
}

/* Function that takes any type that implements the Abser interface as argument. */
func print(a Abser) {
    /* Type of a (Abser) will refer to the type of the variable passed */
    fmt.Printf("Type is %T, value is %f\n", a, a.Abs())
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    fmt.Println(a.Abs())
    print(a)
    print(f)

    a = &v // a *Vertex implements Abser
    fmt.Println(a.Abs())
    print(a)
    print(&v)
    /*
    * Interface Abser doesn't implement Scale()
    * a.Scale(10.0)
    */

    var s Scaler
    s = &v
    s.Scale(10.0)
    fmt.Println(s.Abs())

    l := Line{
       Vertex{16, 12},
       Vertex{4,3},
    }

    a = l /* Line implemets Abser */
    fmt.Println(a.Abs())
    print(a)
    print(l)

    /* Type assertion.
     * If interface `a` is of type `Line`,
     * `t` will have the value and `ok` will be true.
     * else `ok` will be false and `t` will be 0/nil
     */
    t, ok := a.(Line)
    fmt.Println(t, ok)
}
