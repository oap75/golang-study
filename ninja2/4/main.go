package main

import "fmt"

func main() {
	var x = 4
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	b := x << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

}
