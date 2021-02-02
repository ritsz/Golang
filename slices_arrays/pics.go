package main

import "fmt"

func Pics(dx, dy int) {
    picture := make([][]uint8, dy)

    for i := 0; i < len(picture); i++ {
        picture[i] = make([]uint8, dx)
        for j := 0; j < len(picture[i]); j++ {
            picture[i][j] = uint8(i*j)
        }
    }
    for i := 0; i < len(picture); i++ {
        for j := 0; j < len(picture[i]); j++ {
            fmt.Println(picture[i][j])
        }
    }
}
