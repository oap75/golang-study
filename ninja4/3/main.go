package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:5])
	fmt.Println(y)
	z := append(x[5:])
	fmt.Println(z)
	a := append(x[2:7])
	fmt.Println(a)
	b := append(x[1:6])
	fmt.Println(b)
}
