package main

import "unsafe"


func main() {
	println(convert)
}

func convert(p *int) {
	q := (*int32)(unsafe.Pointer(p))
	*q = 0
}
