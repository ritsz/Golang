package main

import (
    "fmt"
)

func main() {
    var cards [3]string
    
    cards = getCards()

    for index, card := range cards {
        fmt.Println(index, card)
    }

    /* Array of arrays - tictactoe */
    tictactoe()
}

func getCards() [3]string {
    cards := [3]string {
        "Ace of Diamonds",
        "Kings of Hearts",
        "Queen of Spades",
    }

    return cards
}
