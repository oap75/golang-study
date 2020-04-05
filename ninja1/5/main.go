package main

import "fmt"

type omid int64

var x omid

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%T", y)
	fmt.Println(y)

}
