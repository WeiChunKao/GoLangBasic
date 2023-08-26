package main

var n int

func main() {
	println(addr())
}

//go:noinline
func addr() (p *int) {
	return &n
}