package main

import "fmt"

func main() {
	var x = struct {
		doors int64
		color string
	}{
		doors: 4,
		color: "blue",
	}
	fmt.Println(x)
}
