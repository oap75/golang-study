package main

import "fmt"

var y = "help Bros"

func main() {
	fmt.Println(y)
	bs := []byte(y)
	fmt.Println(bs)
	rs := []rune(y)
	fmt.Println(rs)
	fmt.Printf("%T\t%#U\n", bs, bs)
	fmt.Printf("%T\t%#U\n", rs, rs)
	y = []byte(y)
	fmt.Println(y)

}
