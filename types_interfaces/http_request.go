package main

import (
	"fmt"
	_ "io"
	"net/http"
	_ "os"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)

	body := resp.Body
	fmt.Printf("%T, %v\n", body, body)
	//
	//  io.Copy(dst, src) copies data from a src that implements Reader interface to a
	//  dst that implements Writer interface.
	//    fmt.Printf("\n**\n")
	//    io.Copy(os.Stdout, body)
	//    fmt.Printf("\n**\n")
	//
	var buf = make([]byte, 200)
	n, _ := body.Read(buf)
	buf = buf[:n]

	fmt.Println(n, string(buf))
}
