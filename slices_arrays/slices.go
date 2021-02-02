package main

import (
    "fmt"
)

/*
 * foobar is of type struct{int, bool}
 * The foobar keyword for this struct can be used when doing initializations
 */
type foobar []struct {
    i int
    b bool
}

/*
 * On this type foobar, implement print function.
 * foobar is called the receiver.
 * the parameter f of receiver type is analogous to this pointer in OOPs.
 */
func (f foobar) print() {
    for _, elem := range f {
        fmt.Println(elem.i, elem.b)
    }
}

func main() {
    var cards []string
    
    /* The foobar keyword for this struct can be used when doing initializations. */
    var s foobar

    /* Slices of struct */
    s = foobar {
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    s.print()
    
    tictactoe_slice()

    cards = getCards()
    cards = append(cards, "Jack of Clubs")

    for index, card := range cards {
        fmt.Println(index, card)
    }

    /*
    *  A slice does not store any data, it just describes a section of an underlying array.
    *  Changing the elements of a slice modifies the corresponding elements of its underlying array.
    * Other slices that share the same underlying array will see those changes.  
    */
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)

    reslice()

    Pics(3,3)
}

func getCards() []string {
    cards := []string {
        "Ace of Diamonds",
        "Kings of Hearts",
        "Queen of Spades",
    }

    return cards
}

func reslice() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length. Notice both LHS and RHS use `s`
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
