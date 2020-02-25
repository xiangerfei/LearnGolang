package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("==============整型==============")
	a := 12
	fmt.Println("Length of a: ", unsafe.Sizeof(a))
	var b int = 12
	fmt.Println("Length of b: ", unsafe.Sizeof(b))
	var c int8 = 12
	fmt.Println("Length of c: ", unsafe.Sizeof(c))
	var d int16 = 12
	fmt.Println("Length of d: ", unsafe.Sizeof(d))
	var e int32 = 12
	fmt.Println("Length of e: ", unsafe.Sizeof(e))
	var f int64 = 12
	fmt.Println("Length of f: ", unsafe.Sizeof(f))
	fmt.Println("==============浮点型==============")
	var boola bool
	fmt.Println("Length of boola: ", unsafe.Sizeof(boola))

}
