package main

import "fmt"

var y = 42

func main() {
	n := fmt.Sprintf("%T\t%v\n", y, y)
	fmt.Printf("%T\t%v\n", n, n)

}
