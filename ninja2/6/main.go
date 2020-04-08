package main

import "fmt"

const (
	a = 2012 + iota
	b = 2012 + iota
	c = 2012 + iota
	d = 2012 + iota
)

func main() {
	fmt.Println(a, b, c, d)

}
