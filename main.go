package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int32
	B []int32
	C string
	D bool
	E struct{}
}

func main() {
	fmt.Println(unsafe.Sizeof(User{}))
	fmt.Println(unsafe.Alignof(User{}))
}
