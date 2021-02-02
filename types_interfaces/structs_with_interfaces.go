package main

import "fmt"

/*
 * Define an interface that any type that wants to be read needs to implement.
 */
type Readable interface {
    Read() []uint8
}

/*
 * ReadBuffer is a struct that has buffer size and any buffer that is Readable ie,
 * implements the Readable interface.
 */
type ReadBuffer struct {
    size int
    buffer Readable
}

/*
 * Define an interface that any type that wants to write needs to implement.
 */
type Writable interface {
    Write([]uint8)
}

type WriteBuffer struct {
    size int
    buffer Writable
}

/*
 * SliceBuffer implements both Readble and Writable interfaces.
 */
type SliceBuffer struct {
    buffer []uint8 
}

func (s SliceBuffer) Read() []uint8 {
    return s.buffer
}

func (s SliceBuffer) Write(in []uint8) {
    s.buffer = in
    fmt.Println("New : ", s.buffer)
}

func main() {
    /* Create a slice buffer */
    buf := SliceBuffer{}
    buf.buffer = []uint8{1, 1, 2, 3, 5}
    fmt.Printf("%T : %v\n", buf, buf)

    /* SliceBuffer is implementing Readable interface, so can be used in ReadBuffer */
    rBuf := ReadBuffer{}
    rBuf.size = len(buf.buffer)
    rBuf.buffer = buf
    fmt.Printf("%T : %v\n", rBuf, rBuf)
    fmt.Println(rBuf.buffer.Read())

    /* SliceBuffer is implementing Writable interface, so can be used in WriteBuffer */
    wBuf := WriteBuffer{}
    wBuf.size = len(buf.buffer)
    wBuf.buffer = buf
    fmt.Printf("%T : %v\n", wBuf, wBuf)
    wBuf.buffer.Write([]uint8{8, 13, 21, 34, 55})
}
