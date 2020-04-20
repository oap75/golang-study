package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
	value  float64
}

type secretAg struct {
	ll  person
	ptk bool
}

func main() {
	a := secretAg{
		ll: person{
			first:  "ali",
			second: "hassan",
			age:    22,
			value:  12.5,
		},
		ptk: true,
	}

	fmt.Println(a.ptk, a.ll.second)
}
