package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%T\t%T\t%T\n", x, y, z)
	fmt.Println(s)
	fmt.Printf("%T\n", s)

}
