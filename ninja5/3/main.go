package main

import "fmt"

type vehicle struct {
	doors int64
	color string
}

type sadan struct {
	vehicle
	loxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 1,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sadan{
		vehicle: vehicle{
			doors: 4,
			color: "yellow",
		},
		loxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)

}
